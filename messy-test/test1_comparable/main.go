package main

import "fmt"

func main() {
	//匿名结构体
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	/*
		sn3 := struct {
			name string
			age  int
		}{age: 11, name: "qq"}
		fmt.Println("sn1==sn3 ? ", sn1==sn3) //编译失败， sn1==sn3 (mismatched types struct {...} and struct {...})
	*/

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		//m   map[string]string
	}{age: 11 /*m: map[string]string{"a": "1"}*/}

	sm2 := struct {
		age int
		//m   map[string]string
	}{age: 11 /*m: map[string]string{"a": "1"}*/}

	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}
	//相同类型的结构体才能够进行比较，结构体是否相同不但与属性类型有关，还与属性顺序相关（sn3和sn1,2就是不可比）
	//能用==比较的前提是结构体内所有的属性类型都是可比较的
	//那什么是可比较的呢？常见的有 bool、数值型、字符、指针、数组等
	//像切片、map、函数等是不能比较的（channel是可比较的）参考 https://golang.org/ref/spec#Comparison_operators
	/*
		m := make(map[string]int)
		n := make(map[string]int)
		fmt.Println(m == nil) //这个是可以的
		fmt.Println(m == n)  //Invalid operation: m == n (operator == is not defined on map[string]int)
	*/
}
