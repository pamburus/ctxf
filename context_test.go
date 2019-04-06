package ctxf

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestContextCreation(t *testing.T) {
	ctx := New(context.Background(), Bool("bool", true), Int("int", 123))
	fields := ctx.Fields()
	assert.Equal(t, 2, cap(fields))
	require.Len(t, fields, 2)
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
