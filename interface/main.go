package main

import "fmt"

type Shape interface {
	Area() float32
}

type Iname interface {
	Mname()
}

type St1 struct{}

func (St1) Mname() {}

type St2 struct{}

func (St2) Mname() {}

func main() {
	//##1.静态类型，动态类型
	var s Shape
	fmt.Println("value of s is", s)    //value of s is <nil>
	fmt.Printf("type of s is %T\n", s) //type of s is <nil>

	//接口的静态类似就是接口本身，而动态类型是实现该接口的类型，也就是具体的实现类的类型
	var i Iname = St1{}            //静态类型Iname，动态类型不固定
	fmt.Printf("type is %T\n", i)  //type is main.St1
	fmt.Printf("value is %v\n", i) //value is {}
	i = St2{}
	fmt.Printf("type is %T\n", i)  //type is main.St2
	fmt.Printf("value is %v\n", i) //value is {}

	//##.2 nil接口值
	var t *St1
	if t == nil {
		fmt.Println("t is nil") //t is nil
	} else {
		fmt.Println("t is not nil")
	}

	//当且仅当动态值（实际分配的值）和动态类型都为 nil 时，接口类型值才为 nil
	var i2 Iname = t       //动态值nil，动态类型*St1
	fmt.Printf("%T\n", i2) //*main.St1
	if i2 == nil {
		fmt.Println("i2 is nil")
	} else {
		fmt.Println("i2 is not nil") //i2 is not nil
	}
	fmt.Printf("i2 is nil pointer:%v", i2 == (*St1)(nil)) //i2 is nil pointer:true

	var inter interface{}
	if inter == nil {
		fmt.Println("nil") //nil
		return
	}
	fmt.Println("not nil")
}
