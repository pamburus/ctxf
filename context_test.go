package ctxf

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockContext struct {
	DeadlineFunc func() (deadline time.Time, ok bool)
	DoneFunc     func() <-chan struct{}
	ErrFunc      func() error
	ValueFunc    func(key interface{}) interface{}
}

func (c mockContext) Deadline() (deadline time.Time, ok bool) {
	return c.DeadlineFunc()
}

func (c mockContext) Done() <-chan struct{} {
	return c.DoneFunc()
}

func (c mockContext) Err() error {
	return c.ErrFunc()
}

func (c mockContext) Value(key interface{}) interface{} {
	return c.ValueFunc(key)
}

func TestContextCreation(t *testing.T) {
	ctx := New(context.Background(), Bool("bool", true), Int("int", 123))
	fields := ctx.Fields()
	assert.Equal(t, 2, cap(fields))
	require.Len(t, fields, 2)
}

func TestContextDeadline(t *testing.T) {
	expected := time.Now().Add(time.Second)
	c := mockContext{
		DeadlineFunc: func() (deadline time.Time, ok bool) {
			return expected, true
		},
	}

	ctx := New(c, Bool("bool", true), Int("int", 123))
	actual, ok := ctx.Deadline()
	assert.Equal(t, true, ok)
	assert.Equal(t, expected, actual)
}

func TestContextDone(t *testing.T) {
	expected := make(<-chan struct{})
	c := mockContext{
		DoneFunc: func() <-chan struct{} {
			return expected
		},
	}

	ctx := New(c, Bool("bool", true), Int("int", 123))
	actual := ctx.Done()
	assert.Equal(t, expected, actual)
}

func TestContextErr(t *testing.T) {
	expected := errors.New("test error")
	c := mockContext{
		ErrFunc: func() error {
			return expected
		},
	}

	ctx := New(c, Bool("bool", true), Int("int", 123))
	actual := ctx.Err()
	assert.Equal(t, expected, actual)
}

func TestContextValue(t *testing.T) {
	goldenKey := 7
	goldenValue := "some"
	c := mockContext{
		ValueFunc: func(key interface{}) interface{} {
			assert.Equal(t, goldenKey, key)

			return goldenValue
		},
	}

	fields := []Field{Bool("bool", true)}
	ctx := New(c, fields...)
	actual := ctx.Value(goldenKey)
	assert.Equal(t, goldenValue, actual)
	assert.Equal(t, fields, ctx.Value(key{}))
}

func TestContextWith(t *testing.T) {
	ctx := New(context.Background())
	fields := []Field{Bool("bool", true)}
	ctx = ctx.With(fields...)
	assert.Equal(t, fields, ctx.Value(key{}))
	fields = append(fields, Int("int", 42))
	ctx = ctx.With(Int("int", 42))
	assert.Equal(t, fields, ctx.Value(key{}))
}

func TestContextWithDeadline(t *testing.T) {
	deadline := time.Now().Add(time.Second)
	ctx := New(context.Background(), Bool("bool", true))
	ctx, cancel := ctx.WithDeadline(deadline)
	defer cancel()
	d, ok := ctx.Deadline()
	assert.Equal(t, true, ok)
	assert.Equal(t, deadline, d)
}

func TestContextWithTimeout(t *testing.T) {
	deadline := time.Now().Add(time.Second)
	ctx := New(context.Background(), Bool("bool", true))
	ctx, cancel := ctx.WithTimeout(time.Second)
	defer cancel()
	d, ok := ctx.Deadline()
	assert.Equal(t, true, ok)
	assert.Equal(t, true, d == deadline || d.After(deadline))
}

func BenchmarkContextCreation(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i != b.N; i++ {
		_ = New(context.Background(), Bool("bool", true), Int("int", 123))
	}
}

func BenchmarkContextAppending(b *testing.B) {
	ctx := New(context.Background(), Bool("bool-1", true), Int("int-1", 123))
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i != b.N; i++ {
		_ = ctx.With(Bool("bool-2", false), Int("int-2", 456))
	}
}
func BenchmarkContextDecoding(b *testing.B) {
	ctx, cancel := context.WithCancel(New(context.Background(), Bool("bool", true), Int("int", 123)))
	defer cancel()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i != b.N; i++ {
		_, _ = Decode(ctx)
	}
}
