package pluginbench

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

const src = `package plugin

func Sum(a int32, b int32) int32 {
	return a + b
}
`

func BenchmarkYaegi(b *testing.B) {
	yi := interp.New(interp.Options{})

	yi.Use(stdlib.Symbols)

	_, err := yi.Eval(src)
	assert.NoError(b, err)

	v, err := yi.Eval("plugin.Sum")
	assert.NoError(b, err)

	sum := v.Interface().(func(int32, int32) int32)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sum(5, 37)
	}
}
