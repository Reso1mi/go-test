package main

import (
	"fmt"
	"go-test/messy-test/test4_init/depend"
)

//一个源文件可以有多个同名的 init函数，执行顺序从上到下，同一个包不同源文件则不一定
// https://golang.org/ref/spec#Package_initialization
func init() {
	fmt.Println("2")
}

func init() {
	fmt.Println("3")
}

func main() {
	depend.GetName()
}
