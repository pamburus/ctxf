# ctxf

[![GoDoc](https://godoc.org/github.com/pamburus/ctxf?status.svg)](https://godoc.org/github.com/pamburus/ctxf)
[![Build Status](https://travis-ci.org/pamburus/ctxf.svg?branch=master)](https://travis-ci.org/pamburus/ctxf)
[![Go Report Status](https://goreportcard.com/badge/github.com/pamburus/ctxf)](https://goreportcard.com/report/github.com/pamburus/ctxf)
[![Coverage Status](https://coveralls.io/repos/github/pamburus/ctxf/badge.svg?branch=master&service=github)](https://coveralls.io/github/pamburus/ctxf?branch=master)

This package provides means for dealing with fields containing of string and valf.Value, which are typified values with snapshotting capabilities uniformly and a way to get its type and value back using visitor pattern.

This package was originally a part of github.com/ssgreg/logf package.

It has become a separate package to get richer flexibility and make this technology available for any other purpose which may not be related to logging.

## Example

The following example creates a new `ctxf` field and gets its value type back using visitor pattern.

```go
package main

import (
	"fmt"

	"github.com/pamburus/ctxf"
)

type testVisitor struct {
    key string
	valf.IgnoringVisitor
}

func (v testVisitor) VisitString(value string) {
	fmt.Printf("[%q] string: %#v\n", v.key, value)
}

func (v testVisitor) VisitBytes(value []byte) {
	fmt.Printf("[%q] bytes: %q\n", v.key, string(value))
}

func main() {
	s := ctxf.String("string", "some string value")
	s.Value.AcceptVisitor(testVisitor{key: s.Key})

	bv := []byte("some bytes value")
	b := ctxf.Bytes("bytes", bv)
	b.Value.AcceptVisitor(testVisitor{key: b.Key})

	bv[1] = 'a'
	b.Value.AcceptVisitor(testVisitor{key: b.Key})

	bs := b.Snapshot()
	bv[1] = 'o'
	bs.Value.AcceptVisitor(testVisitor{key: bs.Key})
}
```

The output is the following:

```
["string"] string: "some string value"
["bytes"] bytes: "some bytes value"
["bytes"] bytes: "same bytes value"
["bytes"] bytes: "same bytes value"
```

