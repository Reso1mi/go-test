package main

import (
	"fmt"
)

type MyInt1 int   //新类型
type MyInt2 = int //别名

func main() {
	var i int = 0
	//var i1 MyInt1 = i 强类型语言需要手动的转换
	var i1 MyInt1 = MyInt1(1)
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
	var cache = make(map[string][]string)
	put(cache)
	fmt.Println(cache["A"])

	//var slice = []string{"A", "B"}
	var b = *new([]int)
	fmt.Println(b == nil)
}

func put(cache map[string][]string) {
	if cache["A"] == nil {
		fmt.Println("ASDSADSADASDAS")
	}
	var temp = []string{"ddddd", "dddddd", "dddddd"}
	cache["A"] = temp
	if cache["A"] == nil {
		fmt.Println("111111111111")
	}
}
