package pluginbench

import (
	"testing"
)

func sumInline(a int32, b int32) int32 {
	return a + b
}

//go:noinline
func sumNoInline(a int32, b int32) int32 {
	return a + b
}

func BenchmarkGoNativeFuncInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumInline(5, 37)
	}
}

func BenchmarkGoNativeFuncNoInline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumNoInline(5, 37)
	}
}
