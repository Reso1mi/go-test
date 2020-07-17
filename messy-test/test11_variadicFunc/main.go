package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(toFullname("carl", "sagan"))
	// output: "carl sagan"
	fmt.Println(toFullname("carl"))
	// output: "carl"
	fmt.Println(toFullname())
	// output: ""

	//Cannot use '[]string{"A", "B", "C"}' (type []string) as type string
	//fmt.Println(toFullname([]string{"A", "B", "C"})) 直接传数组会被当成一个整体，需要使用...解构
	fmt.Println(toFullname([]string{"A", "B", "C"}...))

	//可变参数实际上就是一个切片的语法糖，底层会生成一个切片去append
	fmt.Println(name("imlgw")) //[imlgw resolmi]

	//但是上面的说法也有特例，如果我把现有的切片传递进去就不会生成新切片了
	//会直接使用原切片，所以在chName中可以改变原切片的值，这一点需要注意
	names := []string{"Jack", "Mark", "Alice"}
	chName(names...)
	fmt.Println(names) //[Hacker Mark Alice]
}

//Can only use '...' as final argument in list
//可变参数只能放在最后一个，Java中也是
//func test(id int, names ...string, age int) string {}

func name(names ...string) []string {
	return append(names, "resolmi")
}

func chName(names ...string) {
	names[0] = "Hacker"
}

func toFullname(names ...string) string {
	return strings.Join(names, " ")
}
