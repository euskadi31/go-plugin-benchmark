package pluginbench

import (
	"plugin"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoNativePlugin(t *testing.T) {
	plug, err := plugin.Open("./testdata/goplugin/sum.so")
	assert.NoError(t, err)

	symSum, err := plug.Lookup("Sum")
	assert.NoError(t, err)

	sum := symSum.(func(int32, int32) int32)

	assert.Equal(t, int32(42), sum(5, 37))
}

func BenchmarkGoNativePlugin(b *testing.B) {
	plug, err := plugin.Open("./testdata/goplugin/sum.so")
	assert.NoError(b, err)

	symSum, err := plug.Lookup("Sum")
	assert.NoError(b, err)

	sum := symSum.(func(int32, int32) int32)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sum(5, 37)
	}
}
