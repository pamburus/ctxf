package ctxf

import (
	"context"
	"time"

	"github.com/pamburus/valf"
)

// Context wraps standard context with fast way of adding new fields to it
// and getting all previously added fields.
type Context struct {
	parent context.Context
	fields []Field
}

// Deadline delegates the call to the context.Context.
func (c Context) Deadline() (deadline time.Time, ok bool) {
	return c.parent.Deadline()
}

// Done delegates the call to the context.Context.
func (c Context) Done() <-chan struct{} {
	return c.parent.Done()
}

// Err delegates the call to the context.Context.
func (c Context) Err() error {
	return c.parent.Err()
}

// Value delegates the call to the context.Context.
func (c Context) Value(k interface{}) interface{} {
	if c.fields != nil {
		_, ok := k.(key)
		if ok {
			return c.fields
		}
	}

	return c.parent.Value(k)
}

// With returns a new Context with provided fields appended to its fields.
func (c Context) With(fields ...Field) Context {
	snapshot(fields)

	f := c.fields
	if len(f) == 0 {
		f = fields
	} else {
		f = append(f, fields...)
	}

	c.fields = f[0:len(f):len(f)]

	return c
}

// Fields returns fields associated with the context.
func (c Context) Fields() []Field {
	return c.fields
}

// WithDeadline returns a copy of the c with the deadline adjusted
// to be no later than deadline. If the parent's deadline is already earlier than deadline,
// WithDeadline(deadline) is semantically equivalent to c. The returned
// context's Done channel is closed when the deadline expires, when the returned
// cancel function is called, or when the parent context's Done channel is
// closed, whichever happens first.
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this Context complete.
func (c Context) WithDeadline(deadline time.Time) (Context, context.CancelFunc) {
	var cancel context.CancelFunc
	c.parent, cancel = context.WithDeadline(c.parent, deadline)

	return c, cancel
}

// WithTimeout returns c.WithDeadline(time.Now().Add(timeout)).
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this Context complete.
func (c Context) WithTimeout(timeout time.Duration) (Context, context.CancelFunc) {
	var cancel context.CancelFunc
	c.parent, cancel = context.WithTimeout(c.parent, timeout)

	return c, cancel
}

// WithCancel returns a copy of c with a new Done channel. The returned
// context's Done channel is closed when the returned cancel function is called
// or when the parent context's Done channel is closed, whichever happens first.
//
// Canceling this context releases resources associated with it, so code should
// call cancel as soon as the operations running in this Context complete.
func (c Context) WithCancel() (Context, context.CancelFunc) {
	var cancel context.CancelFunc
	c.parent, cancel = context.WithCancel(c.parent)

	return c, cancel
}

// WithValue returns a copy of Context in which the value associated with key is
// value.
func (c Context) WithValue(key, value interface{}) Context {
	c.parent = context.WithValue(c.parent, key, value)

	return c
}

// Encode casts Context to standard context.Context.
func (c Context) Encode() context.Context {
	return c
}

// New returns a new Context with provided fields appended to it.
func New(parent context.Context, fields ...Field) Context {
	snapshot(fields)

	return Context{parent, fields}
}

// Fields returns all fields from context previously added to it with New.
func Fields(ctx context.Context) []Field {
	c, ok := Decode(ctx)
	if !ok {
		return nil
	}

	return c.fields
}

// Decode retreives fields associated with the ctx and returns Context
// with that fields and a flag which indicates where the ctx was constructed with New.
func Decode(ctx context.Context) (Context, bool) {
	switch c := ctx.(type) {
	case *Context:
		return *c, true
	case Context:
		return c, true
	default:
		value := ctx.Value(key{})
		if value == nil {
			return Context{ctx, nil}, false
		}

		return Context{ctx, value.([]Field)}, true
	}
}

// DecodeOptional retreives fields associated with the ctx and returns Context
// with that fields.
func DecodeOptional(ctx context.Context) Context {
	c, _ := Decode(ctx)

	return c
}

type key struct{}

func snapshot(fields []Field) {
	for i := range fields {
		valf.Snapshot(&fields[i].Value)
	}
}
