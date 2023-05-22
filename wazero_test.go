package pluginbench

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

func TestWazeroRustSum(t *testing.T) {
	wasmBytes, err := os.ReadFile("./testdata/rust-sum/target/wasm32-unknown-unknown/release/rust_sum.wasm")
	assert.NoError(t, err)

	ctx := context.Background()

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx) // This closes everything this Runtime created.

	// Instantiate WASI, which implements host functions needed for TinyGo to
	// implement `panic`.
	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	// Instantiate the guest Wasm into the same runtime. It exports the `add`
	// function, implemented in WebAssembly.
	mod, err := r.Instantiate(ctx, wasmBytes)
	assert.NoError(t, err)

	sum := mod.ExportedFunction("sum")

	v, err := sum.Call(ctx, 5, 37)
	assert.NoError(t, err)

	assert.Equal(t, int32(42), int32(v[0]))
}

func TestWazeroGoSum(t *testing.T) {
	wasmBytes, err := os.ReadFile("./testdata/go-sum/go_sum.wasm")
	assert.NoError(t, err)

	ctx := context.Background()

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx) // This closes everything this Runtime created.

	// Instantiate WASI, which implements host functions needed for TinyGo to
	// implement `panic`.
	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	// Instantiate the guest Wasm into the same runtime. It exports the `add`
	// function, implemented in WebAssembly.
	mod, err := r.Instantiate(ctx, wasmBytes)
	assert.NoError(t, err)

	sum := mod.ExportedFunction("Sum")

	v, err := sum.Call(ctx, 5, 37)
	assert.NoError(t, err)

	assert.Equal(t, int32(42), int32(v[0]))
}

func BenchmarkWazeroRustSum(b *testing.B) {
	wasmBytes, err := os.ReadFile("./testdata/rust-sum/target/wasm32-unknown-unknown/release/rust_sum.wasm")
	assert.NoError(b, err)

	ctx := context.Background()

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx) // This closes everything this Runtime created.

	// Instantiate WASI, which implements host functions needed for TinyGo to
	// implement `panic`.
	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	// Instantiate the guest Wasm into the same runtime. It exports the `add`
	// function, implemented in WebAssembly.
	mod, err := r.Instantiate(ctx, wasmBytes)
	assert.NoError(b, err)

	sum := mod.ExportedFunction("sum")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sum.Call(ctx, 5, 37)
	}
}

func BenchmarkWazeroGoSum(b *testing.B) {
	wasmBytes, err := os.ReadFile("./testdata/go-sum/go_sum.wasm")
	assert.NoError(b, err)

	ctx := context.Background()

	// Create a new WebAssembly Runtime.
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx) // This closes everything this Runtime created.

	// Instantiate WASI, which implements host functions needed for TinyGo to
	// implement `panic`.
	wasi_snapshot_preview1.MustInstantiate(ctx, r)

	// Instantiate the guest Wasm into the same runtime. It exports the `add`
	// function, implemented in WebAssembly.
	mod, err := r.Instantiate(ctx, wasmBytes)
	assert.NoError(b, err)

	sum := mod.ExportedFunction("Sum")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sum.Call(ctx, 5, 37)
	}
}
