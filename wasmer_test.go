package pluginbench

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wasmerio/wasmer-go/wasmer"
)

func TestWasmerRustSum(t *testing.T) {
	wasmBytes, err := os.ReadFile("./testdata/rust-sum/target/wasm32-unknown-unknown/release/rust_sum.wasm")
	assert.NoError(t, err)

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	module, err := wasmer.NewModule(store, wasmBytes)
	assert.NoError(t, err)

	// Instantiates the module
	importObject := wasmer.NewImportObject()
	instance, err := wasmer.NewInstance(module, importObject)
	assert.NoError(t, err)
	defer instance.Close()

	// Gets the `sum` exported function from the WebAssembly instance.
	sum, err := instance.Exports.GetFunction("sum")
	assert.NoError(t, err)

	v, err := sum(5, 37)
	assert.NoError(t, err)

	assert.Equal(t, int32(42), v.(int32))
}

func TestWasmerGoSum(t *testing.T) {
	wasmBytes, err := os.ReadFile("./testdata/go-sum/go_sum.wasm")
	assert.NoError(t, err)

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	module, err := wasmer.NewModule(store, wasmBytes)
	assert.NoError(t, err)

	// Instantiates the module
	importObject := wasmer.NewImportObject()
	instance, err := wasmer.NewInstance(module, importObject)
	assert.NoError(t, err)
	defer instance.Close()

	// Gets the `sum` exported function from the WebAssembly instance.
	sum, err := instance.Exports.GetFunction("Sum")
	assert.NoError(t, err)

	v, err := sum(5, 37)
	assert.NoError(t, err)

	assert.Equal(t, int32(42), v.(int32))
}

func BenchmarkWasmerRustSum(b *testing.B) {
	wasmBytes, err := os.ReadFile("./testdata/rust-sum/target/wasm32-unknown-unknown/release/rust_sum.wasm")
	assert.NoError(b, err)

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	module, err := wasmer.NewModule(store, wasmBytes)
	assert.NoError(b, err)

	// Instantiates the module
	importObject := wasmer.NewImportObject()
	instance, err := wasmer.NewInstance(module, importObject)
	assert.NoError(b, err)
	defer instance.Close()

	// Gets the `sum` exported function from the WebAssembly instance.
	sum, err := instance.Exports.GetFunction("sum")
	assert.NoError(b, err)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sum(5, 37)
	}
}

func BenchmarkWasmerGoSum(b *testing.B) {
	wasmBytes, err := os.ReadFile("./testdata/go-sum/go_sum.wasm")
	assert.NoError(b, err)

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	module, err := wasmer.NewModule(store, wasmBytes)
	assert.NoError(b, err)

	// Instantiates the module
	importObject := wasmer.NewImportObject()
	instance, err := wasmer.NewInstance(module, importObject)
	assert.NoError(b, err)
	defer instance.Close()

	// Gets the `sum` exported function from the WebAssembly instance.
	sum, err := instance.Exports.GetFunction("Sum")
	assert.NoError(b, err)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sum(5, 37)
	}
}
