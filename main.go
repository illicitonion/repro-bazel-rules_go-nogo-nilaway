package main

import (
	"fmt"
	"example/mylib"
)

func main() {
	//nolint:nilaway
	s := mylib.MakeANilPanicStruct()
	//nolint:nilaway
	s.OhnoIMayPanic()

	fmt.Println("Hello, World!")
}
