package main

import "fmt"

type User struct {
	name string
	age  int
}

func (u *User) getName() string {
	return u.name
}

func (u User) getAge() int {
	return u.age
}

type N int

func (n N) test() {
	fmt.Printf("test n: %p, %v\n", &n, n)
}

func main() {
	u := &User{"imlgw", 12}
	//##1. 通过类型引用的method expression会被还原成普通函数样式
	f1 := (*User).getName
	fmt.Println(f1(u)) //需要手动传递传递receiver
	//f := User.getName //invalid method expression User.getName (needs pointer receiver: (*User).getName)
	//fmt.Println(f(*u))
	f2 := User.getAge
	fmt.Println(f2(*u)) //12

	ValueType()
	PointerType()
}

func PointerType() {
	//##2. 通过实例或指针引用的method value函数签名不变，依然按照原来的方式调用
	var n N = 100
	var p = &n

	n++
	//这一步会直接复制f3执行需要的receiver对象
	f3 := n.test //test的receiver是值，所以这里直接复制n等于101

	n++
	f4 := p.test //复制*p(将p中的值取出来复制一份) 等于102

	n++
	fmt.Printf("main.n : %p, %v\n", &n, n) //main.n : 0xc0000160e8, 103
	f3()                                   //test n: 0xc0000160f8, 101
	f4()                                   //test n: 0xc000016118, 102
}

func ValueType() {
	u := &User{"imlgw", 12}
	//##1. 通过类型引用的method expression会被还原成普通函数样式
	f1 := (*User).getName
	fmt.Println(f1(u)) //需要手动传递传递receiver
	//f := User.getName //invalid method expression User.getName (needs pointer receiver: (*User).getName)
	//fmt.Println(f(*u))
	f2 := User.getAge
	fmt.Println(f2(*u)) //12

	//##2. 通过实例或指针引用的method value函数签名不变，依然按照原来的方式调用
	var n N = 100
	var p = &n

	n++
	//这一步会直接复制f3执行需要的receiver对象
	f3 := n.test //test的receiver是值，所以这里直接复制n等于101

	n++
	f4 := p.test //复制*p(将p中的值取出来复制一份) 等于102

	n++
	fmt.Printf("main.n : %p, %v\n", &n, n) //main.n : 0xc0000160e8, 103
	f3()                                   //test n: 0xc0000160f8, 101
	f4()                                   //test n: 0xc000016118, 102
}
