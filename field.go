package ctxf

import (
	"fmt"
	"time"

	"github.com/pamburus/valf"
)

// Field is a key and value pair.
type Field struct {
	Key   string
	Value valf.Value
}

// Snapshot returns a Field which can be safely stored for a long with guarantee that it won't be modified.
// The data of the field are copied if needed to achieve that guarantee.
func (f Field) Snapshot() Field {
	f.Value = f.Value.Snapshot()

	return f
}

// Bool returns a new Field with the given key and bool.
func Bool(k string, v bool) Field {
	return Field{Key: k, Value: valf.Bool(v)}
}

// Int returns a new Field with the given key and int.
func Int(k string, v int) Field {
	return Field{Key: k, Value: valf.Int(v)}
}

// Int64 returns a new Field with the given key and int64.
func Int64(k string, v int64) Field {
	return Field{Key: k, Value: valf.Int64(v)}
}

// Int32 returns a new Field with the given key and int32.
func Int32(k string, v int32) Field {
	return Field{Key: k, Value: valf.Int32(v)}
}

// Int16 returns a new Field with the given key and int16.
func Int16(k string, v int16) Field {
	return Field{Key: k, Value: valf.Int16(v)}
}

// Int8 returns a new Field with the given key and int.8
func Int8(k string, v int8) Field {
	return Field{Key: k, Value: valf.Int8(v)}
}

// Uint returns a new Field with the given key and uint.
func Uint(k string, v uint) Field {
	return Field{Key: k, Value: valf.Uint(v)}
}

// Uint64 returns a new Field with the given key and uint64.
func Uint64(k string, v uint64) Field {
	return Field{Key: k, Value: valf.Uint64(v)}
}

// Uint32 returns a new Field with the given key and uint32.
func Uint32(k string, v uint32) Field {
	return Field{Key: k, Value: valf.Uint32(v)}
}

// Uint16 returns a new Field with the given key and uint16.
func Uint16(k string, v uint16) Field {
	return Field{Key: k, Value: valf.Uint16(v)}
}

// Uint8 returns a new Field with the given key and uint8.
func Uint8(k string, v uint8) Field {
	return Field{Key: k, Value: valf.Uint8(v)}
}

// Float64 returns a new Field with the given key and float64.
func Float64(k string, v float64) Field {
	return Field{Key: k, Value: valf.Float64(v)}
}

// Float32 returns a new Field with the given key and float32.
func Float32(k string, v float32) Field {
	return Field{Key: k, Value: valf.Float32(v)}
}

// Duration returns a new Field with the given key and time.Duration.
func Duration(k string, v time.Duration) Field {
	return Field{Key: k, Value: valf.Duration(v)}
}

// Bytes returns a new Field with the given key and slice of bytes.
func Bytes(k string, v []byte) Field {
	return Field{Key: k, Value: valf.Bytes(v)}
}

// String returns a new Field with the given key and string.
func String(k string, v string) Field {
	return Field{Key: k, Value: valf.String(v)}
}

// Strings returns a new Field with the given key and slice of strings.
func Strings(k string, v []string) Field {
	return Field{Key: k, Value: valf.Strings(v)}
}

// Bools returns a new Field with the given key and slice of bools.
func Bools(k string, v []bool) Field {
	return Field{Key: k, Value: valf.Bools(v)}
}

// Ints returns a new Field with the given key and slice of ints.
func Ints(k string, v []int) Field {
	return Field{Key: k, Value: valf.Ints(v)}
}

// Ints8 returns a new Field with the given key and slice of 8-bit ints.
func Ints8(k string, v []int8) Field {
	return Field{Key: k, Value: valf.Ints8(v)}
}

// Ints16 returns a new Field with the given key and slice of 16-bit ints.
func Ints16(k string, v []int16) Field {
	return Field{Key: k, Value: valf.Ints16(v)}
}

// Ints32 returns a new Field with the given key and slice of 32-bit ints.
func Ints32(k string, v []int32) Field {
	return Field{Key: k, Value: valf.Ints32(v)}
}

// Ints64 returns a new Field with the given key and slice of 64-bit ints.
func Ints64(k string, v []int64) Field {
	return Field{Key: k, Value: valf.Ints64(v)}
}

// Uints returns a new Field with the given key and slice of uints.
func Uints(k string, v []uint) Field {
	return Field{Key: k, Value: valf.Uints(v)}
}

// Uints8 returns a new Field with the given key and slice of 8-bit uints.
func Uints8(k string, v []uint8) Field {
	return Field{Key: k, Value: valf.Uints8(v)}
}

// Uints16 returns a new Field with the given key and slice of 16-bit uints.
func Uints16(k string, v []uint16) Field {
	return Field{Key: k, Value: valf.Uints16(v)}
}

// Uints32 returns a new Field with the given key and slice of 32-bit uints.
func Uints32(k string, v []uint32) Field {
	return Field{Key: k, Value: valf.Uints32(v)}
}

// Uints64 returns a new Field with the given key and slice of 64-bit uints.
func Uints64(k string, v []uint64) Field {
	return Field{Key: k, Value: valf.Uints64(v)}
}

// Floats32 returns a new Field with the given key and slice of 32-bit floats.
func Floats32(k string, v []float32) Field {
	return Field{Key: k, Value: valf.Floats32(v)}
}

// Floats64 returns a new Field with the given key and slice of 64-biy floats.
func Floats64(k string, v []float64) Field {
	return Field{Key: k, Value: valf.Floats64(v)}
}

// Durations returns a new Field with the given key and slice of time.Duration.
func Durations(k string, v []time.Duration) Field {
	return Field{Key: k, Value: valf.Durations(v)}
}

// ConstBytes returns a new Field with the given key and slice of bytes.
//
// Call ConstBytes if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstBytes(k string, v []byte) Field {
	return Field{Key: k, Value: valf.ConstBytes(v)}
}

// ConstBools returns a new Field with the given key and slice of bools.
//
// Call ConstBools if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstBools(k string, v []bool) Field {
	return Field{Key: k, Value: valf.ConstBools(v)}
}

// ConstInts returns a new Field with the given key and slice of ints.
//
// Call ConstInts if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstInts(k string, v []int) Field {
	return Field{Key: k, Value: valf.ConstInts(v)}
}

// ConstInts8 returns a new Field with the given key and slice of 8-bit ints.
//
// Call ConstInts8 if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstInts8(k string, v []int8) Field {
	return Field{Key: k, Value: valf.ConstInts8(v)}
}

// ConstInts16 returns a new Field with the given key and slice of 16-bit ints.
//
// Call ConstInts16 if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstInts16(k string, v []int16) Field {
	return Field{Key: k, Value: valf.ConstInts16(v)}
}

// ConstInts32 returns a new Field with the given key and slice of 32-bit ints.
//
// Call ConstInts32 if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstInts32(k string, v []int32) Field {
	return Field{Key: k, Value: valf.ConstInts32(v)}
}

// ConstInts64 returns a new Field with the given key and slice of 64-bit ints.
//
// Call ConstInts64 if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstInts64(k string, v []int64) Field {
	return Field{Key: k, Value: valf.ConstInts64(v)}
}

// ConstUints returns a new Field with the given key and slice of uints.
//
// Call ConstUints if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstUints(k string, v []uint) Field {
	return Field{Key: k, Value: valf.ConstUints(v)}
}

// ConstUints8 returns a new Field with the given key and slice of 8-bit uints.
//
// Call ConstUints8 if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstUints8(k string, v []uint8) Field {
	return Field{Key: k, Value: valf.ConstUints8(v)}
}

// ConstUints16 returns a new Field with the given key and slice of 16-bit uints.
//
// Call ConstUints16 if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstUints16(k string, v []uint16) Field {
	return Field{Key: k, Value: valf.ConstUints16(v)}
}

// ConstUints32 returns a new Field with the given key and slice of 32-bit uints.
//
// Call ConstUints32 if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstUints32(k string, v []uint32) Field {
	return Field{Key: k, Value: valf.ConstUints32(v)}
}

// ConstUints64 returns a new Field with the given key and slice of 64-bit uints.
//
// Call ConstUints64 if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstUints64(k string, v []uint64) Field {
	return Field{Key: k, Value: valf.ConstUints64(v)}
}

// ConstFloats32 returns a new Field with the given key and slice of 32-bit floats.
//
// Call ConstFloats32 if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstFloats32(k string, v []float32) Field {
	return Field{Key: k, Value: valf.ConstFloats32(v)}
}

// ConstFloats64 returns a new Field with the given key and slice of 64-bit floats.
//
// Call ConstFloats64 if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstFloats64(k string, v []float64) Field {
	return Field{Key: k, Value: valf.ConstFloats64(v)}
}

// ConstDurations returns a new Field with the given key and slice of time.Duration.
//
// Call ConstDurations if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstDurations(k string, v []time.Duration) Field {
	return Field{Key: k, Value: valf.ConstDurations(v)}
}

// NamedError returns a new Field with the given key and error.
func NamedError(k string, v error) Field {
	return Field{Key: k, Value: valf.Error(v)}
}

// Error returns a new Field with the given error and key "error".
func Error(v error) Field {
	return NamedError("error", v)
}

// Time returns a new Field with the given key and time.Time.
func Time(k string, v time.Time) Field {
	return Field{Key: k, Value: valf.Time(v)}
}

// Array returns a new Field with the given key and valf.ValueArray.
func Array(k string, v valf.ValueArray) Field {
	return Field{Key: k, Value: valf.Array(v)}
}

// ConstArray returns a new Field with the given key and immutable valf.ValueArray.
func ConstArray(k string, v valf.ValueArray) Field {
	return Field{Key: k, Value: valf.ConstArray(v)}
}

// Object returns a new Field with the given key and valf.ValueObject.
func Object(k string, v valf.ValueObject) Field {
	return Field{Key: k, Value: valf.Object(v)}
}

// ConstObject returns a new Field with the given key and immutable valf.ValueArray.
func ConstObject(k string, v valf.ValueObject) Field {
	return Field{Key: k, Value: valf.ConstObject(v)}
}

// ConstStringer returns a new Field with the given key and Stringer.
// Call ConstStringer if your object is const. It has significantly less
// impact on the calling goroutine.
func ConstStringer(k string, v fmt.Stringer) Field {
	return Field{Key: k, Value: valf.ConstStringer(v)}
}

// Stringer returns a new Field with the given key and Stringer.
func Stringer(k string, v fmt.Stringer) Field {
	return Field{Key: k, Value: valf.Stringer(v)}
}

// ConstFormatter returns a new Field with the given key and verb and interface
// to format.
//
// Call ConstFormatter if your object is const. It has significantly less
// impact on the calling goroutine.
//
func ConstFormatter(k string, verb string, v interface{}) Field {
	return Field{Key: k, Value: valf.ConstFormatter(verb, v)}
}

// ConstFormatterRepr returns a new Field with the given key and interface to
// format. It uses the predefined verb "%#v" (a Go-syntax representation of
// the value).
//
// Call ConstFormatterRepr if your object is const. It has significantly less
// impact on the calling goroutine.
//
func ConstFormatterRepr(k string, v interface{}) Field {
	return Field{Key: k, Value: valf.ConstFormatterRepr(v)}
}

// Formatter returns a new Field with the given key and verb and interface to
// format.
func Formatter(k string, verb string, v interface{}) Field {
	return Field{Key: k, Value: valf.Formatter(verb, v)}
}

// FormatterRepr returns a new Field with the given key and interface to format.
// It uses the predefined verb "%#v" (a Go-syntax representation of the value).
func FormatterRepr(k string, v interface{}) Field {
	return Field{Key: k, Value: valf.FormatterRepr(v)}
}

// Any returns a new Field with the given key and value of any type. It tries
// to choose the best way to represent a Value.
//
// Note that Any is not able to choose ConstX methods. Use specific Field
// methods for better performance.
func Any(k string, v interface{}) Field {
	return Field{Key: k, Value: valf.Any(v)}
}

// ConstAny returns a new Field with the given key and value of any type. It tries
// to choose the best way to represent a Value.
func ConstAny(k string, v interface{}) Field {
	return Field{Key: k, Value: valf.ConstAny(v)}
}
