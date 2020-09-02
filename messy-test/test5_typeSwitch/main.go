package main

import (
	"fmt"
)

func GetValue() int {
	return 1
}

func GetInterfaceVal() interface{} {
	return 1
}

func main() {
	//i := GetValue() 只有接口才能进行类型选择
	i := GetInterfaceVal()
	//interface.(type)只能用于switch语句中
	switch i.(type) {
	case int:
		println("int") //int
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
	var n = []int{1, 2, 3, 4, 5, 6}
	m := n[0:3:6]
	fmt.Println(m, len(m), cap(m))
}
