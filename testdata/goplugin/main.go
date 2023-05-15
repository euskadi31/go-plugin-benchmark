package main

//go:generate go build -o sum.so -buildmode=plugin main.go

func main() {

}

func Sum(a int32, b int32) int32 {
	return a + b
}
