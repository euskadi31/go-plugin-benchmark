package main

//go:generate tinygo build -o go_sum.wasm -target wasi -panic=trap -scheduler=none ./main.go
func main() {

}

//export Sum
func Sum(a int32, b int32) int32 {
	return a + b
}
