package ctxf

import (
	"context"
	"testing"
	"time"

	"github.com/pamburus/valf"
	"github.com/stretchr/testify/assert"
)

type customValue struct {
	value int
}

func TestField(t *testing.T) {
	tcs := []struct {
		Name string
		Key  string
		Type valf.Type
		Make func() Field
	}{
		{
			Name: "None",
			Type: valf.TypeNone,
			Make: func() Field {
				return Field{Key: "None"}
			},
		},
		{
			Name: "Any",
			Type: valf.TypeAny,
			Make: func() Field {
				return Any("Any", customValue{42})
			},
		},
		{
			Name: "ConstAny",
			Type: valf.TypeAny,
			Make: func() Field {
				return ConstAny("ConstAny", customValue{42})
			},
		},
		{
			Name: "Bool",
			Type: valf.TypeBool,
			Make: func() Field {
				return Bool("Bool", true)
			},
		},
		{
			Name: "Int",
			Type: valf.TypeInt,
			Make: func() Field {
				return Int("Int", 42)
			},
		},
		{
			Name: "Int8",
			Type: valf.TypeInt8,
			Make: func() Field {
				return Int8("Int8", 42)
			},
		},
		{
			Name: "Int16",
			Type: valf.TypeInt16,
			Make: func() Field {
				return Int16("Int16", 42)
			},
		},
		{
			Name: "Int32",
			Type: valf.TypeInt32,
			Make: func() Field {
				return Int32("Int32", 42)
			},
		},
		{
			Name: "Int64",
			Type: valf.TypeInt64,
			Make: func() Field {
				return Int64("Int64", 42)
			},
		},
		{
			Name: "Uint",
			Type: valf.TypeUint,
			Make: func() Field {
				return Uint("Uint", 42)
			},
		},
		{
			Name: "Uint8",
			Type: valf.TypeUint8,
			Make: func() Field {
				return Uint8("Uint8", 42)
			},
		},
		{
			Name: "Uint16",
			Type: valf.TypeUint16,
			Make: func() Field {
				return Uint16("Uint16", 42)
			},
		},
		{
			Name: "Uint32",
			Type: valf.TypeUint32,
			Make: func() Field {
				return Uint32("Uint32", 42)
			},
		},
		{
			Name: "Uint64",
			Type: valf.TypeUint64,
			Make: func() Field {
				return Uint64("Uint64", 42)
			},
		},
		{
			Name: "Float32",
			Type: valf.TypeFloat32,
			Make: func() Field {
				return Float32("Float32", 0.42)
			},
		},
		{
			Name: "Float64",
			Type: valf.TypeFloat64,
			Make: func() Field {
				return Float64("Float64", 0.42)
			},
		},
		{
			Name: "Duration",
			Type: valf.TypeDuration,
			Make: func() Field {
				return Duration("Duration", time.Second)
			},
		},
		{
			Name: "Error",
			Type: valf.TypeError,
			Key:  "error",
			Make: func() Field {
				return Error(context.Canceled)
			},
		},
		{
			Name: "NamedError",
			Type: valf.TypeError,
			Make: func() Field {
				return NamedError("NamedError", context.Canceled)
			},
		},
		{
			Name: "Time",
			Type: valf.TypeTime,
			Make: func() Field {
				return Time("Time", time.Now())
			},
		},
		{
			Name: "String",
			Type: valf.TypeString,
			Make: func() Field {
				return String("String", "value")
			},
		},
		{
			Name: "Bytes",
			Type: valf.TypeBytes,
			Make: func() Field {
				return Bytes("Bytes", []byte{1, 2})
			},
		},
		{
			Name: "ConstBytes",
			Type: valf.TypeBytes,
			Make: func() Field {
				return ConstBytes("ConstBytes", []byte{1, 2})
			},
		},
		{
			Name: "Bools",
			Type: valf.TypeBools,
			Make: func() Field {
				return Bools("Bools", []bool{false, true})
			},
		},
		{
			Name: "ConstBools",
			Type: valf.TypeBools,
			Make: func() Field {
				return ConstBools("ConstBools", []bool{false, true})
			},
		},
		{
			Name: "Ints",
			Type: valf.TypeInts,
			Make: func() Field {
				return Ints("Ints", []int{21, 42})
			},
		},
		{
			Name: "ConstInts",
			Type: valf.TypeInts,
			Make: func() Field {
				return ConstInts("ConstInts", []int{21, 42})
			},
		},
		{
			Name: "Ints8",
			Type: valf.TypeInts8,
			Make: func() Field {
				return Ints8("Ints8", []int8{21, 42})
			},
		},
		{
			Name: "ConstInts8",
			Type: valf.TypeInts8,
			Make: func() Field {
				return ConstInts8("ConstInts8", []int8{21, 42})
			},
		},
		{
			Name: "Ints16",
			Type: valf.TypeInts16,
			Make: func() Field {
				return Ints16("Ints16", []int16{21, 42})
			},
		},
		{
			Name: "ConstInts16",
			Type: valf.TypeInts16,
			Make: func() Field {
				return ConstInts16("ConstInts16", []int16{21, 42})
			},
		},
		{
			Name: "Ints32",
			Type: valf.TypeInts32,
			Make: func() Field {
				return Ints32("Ints32", []int32{21, 42})
			},
		},
		{
			Name: "ConstInts32",
			Type: valf.TypeInts32,
			Make: func() Field {
				return ConstInts32("ConstInts32", []int32{21, 42})
			},
		},
		{
			Name: "Ints64",
			Type: valf.TypeInts64,
			Make: func() Field {
				return Ints64("Ints64", []int64{21, 42})
			},
		},
		{
			Name: "ConstInts64",
			Type: valf.TypeInts64,
			Make: func() Field {
				return ConstInts64("ConstInts64", []int64{21, 42})
			},
		},
		{
			Name: "Uints",
			Type: valf.TypeUints,
			Make: func() Field {
				return Uints("Uints", []uint{21, 42})
			},
		},
		{
			Name: "ConstUints",
			Type: valf.TypeUints,
			Make: func() Field {
				return ConstUints("ConstUints", []uint{21, 42})
			},
		},
		{
			Name: "Uints8",
			Type: valf.TypeUints8,
			Make: func() Field {
				return Uints8("Uints8", []uint8{21, 42})
			},
		},
		{
			Name: "ConstUints8",
			Type: valf.TypeUints8,
			Make: func() Field {
				return ConstUints8("ConstUints8", []uint8{21, 42})
			},
		},
		{
			Name: "Uints16",
			Type: valf.TypeUints16,
			Make: func() Field {
				return Uints16("Uints16", []uint16{21, 42})
			},
		},
		{
			Name: "ConstUints16",
			Type: valf.TypeUints16,
			Make: func() Field {
				return ConstUints16("ConstUints16", []uint16{21, 42})
			},
		},
		{
			Name: "Uints32",
			Type: valf.TypeUints32,
			Make: func() Field {
				return Uints32("Uints32", []uint32{21, 42})
			},
		},
		{
			Name: "ConstUints32",
			Type: valf.TypeUints32,
			Make: func() Field {
				return ConstUints32("ConstUints32", []uint32{21, 42})
			},
		},
		{
			Name: "Uints64",
			Type: valf.TypeUints64,
			Make: func() Field {
				return Uints64("Uints64", []uint64{21, 42})
			},
		},
		{
			Name: "ConstUints64",
			Type: valf.TypeUints64,
			Make: func() Field {
				return ConstUints64("ConstUints64", []uint64{21, 42})
			},
		},
		{
			Name: "Floats32",
			Type: valf.TypeFloats32,
			Make: func() Field {
				return Floats32("Floats32", []float32{0.21, 0.42})
			},
		},
		{
			Name: "ConstFloats32",
			Type: valf.TypeFloats32,
			Make: func() Field {
				return ConstFloats32("ConstFloats32", []float32{0.21, 0.42})
			},
		},
		{
			Name: "Floats64",
			Type: valf.TypeFloats64,
			Make: func() Field {
				return Floats64("Floats64", []float64{0.21, 0.42})
			},
		},
		{
			Name: "ConstFloats64",
			Type: valf.TypeFloats64,
			Make: func() Field {
				return ConstFloats64("ConstFloats64", []float64{0.21, 0.42})
			},
		},
		{
			Name: "Durations",
			Type: valf.TypeDurations,
			Make: func() Field {
				return Durations("Durations", []time.Duration{time.Second, time.Hour})
			},
		},
		{
			Name: "ConstDurations",
			Type: valf.TypeDurations,
			Make: func() Field {
				return ConstDurations("ConstDurations", []time.Duration{time.Second, time.Hour})
			},
		},
		{
			Name: "Strings",
			Type: valf.TypeStrings,
			Make: func() Field {
				return Strings("Strings", []string{"value-1", "value-2"})
			},
		},
		{
			Name: "ConstStrings",
			Type: valf.TypeStrings,
			Make: func() Field {
				return ConstStrings("ConstStrings", []string{"value-1", "value-2"})
			},
		},
		{
			Name: "Array",
			Type: valf.TypeArray,
			Make: func() Field {
				return Array("Array", nil)
			},
		},
		{
			Name: "ConstArray",
			Type: valf.TypeArray,
			Make: func() Field {
				return ConstArray("ConstArray", nil)
			},
		},
		{
			Name: "Object",
			Type: valf.TypeObject,
			Make: func() Field {
				return Object("Object", nil)
			},
		},
		{
			Name: "ConstObject",
			Type: valf.TypeObject,
			Make: func() Field {
				return ConstObject("ConstObject", nil)
			},
		},
		{
			Name: "Stringer",
			Type: valf.TypeStringer,
			Make: func() Field {
				return Stringer("Stringer", nil)
			},
		},
		{
			Name: "ConstStringer",
			Type: valf.TypeStringer,
			Make: func() Field {
				return ConstStringer("ConstStringer", nil)
			},
		},
		{
			Name: "Formatter",
			Type: valf.TypeFormatter,
			Make: func() Field {
				return Formatter("Formatter", "%v", nil)
			},
		},
		{
			Name: "ConstFormatter",
			Type: valf.TypeFormatter,
			Make: func() Field {
				return ConstFormatter("ConstFormatter", "%v", nil)
			},
		},
		{
			Name: "FormatterRepr",
			Type: valf.TypeFormatter,
			Make: func() Field {
				return FormatterRepr("FormatterRepr", nil)
			},
		},
		{
			Name: "ConstFormatterRepr",
			Type: valf.TypeFormatter,
			Make: func() Field {
				return ConstFormatterRepr("ConstFormatterRepr", nil)
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			field := tc.Make()
			key := tc.Key
			if key == "" {
				key = tc.Name
			}
			assert.Equal(t, tc.Type, field.Value.Type())
			assert.Equal(t, key, field.Key)
		})
	}
}

func TestFieldSnapshot(t *testing.T) {
	key := "key"
	field := Bytes(key, []byte{1, 2})
	assert.Equal(t, key, field.Key)
	assert.Equal(t, false, field.Value.Const())
	field = field.Snapshot()
	assert.Equal(t, key, field.Key)
	assert.Equal(t, true, field.Value.Const())
}
