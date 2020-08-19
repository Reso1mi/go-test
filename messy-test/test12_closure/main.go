package main

import (
	"fmt"
	"time"
)

func foo(i int) func() {
	return func() {
		i++
		fmt.Println(i)
	}
}

//引用传递（本质还是指传递）
func foo1(x *int) func() {
	return func() {
		*x = *x + 1
		fmt.Printf("foo1 val = %d\n", *x)
	}
}

//值传递
func foo2(x int) func() {
	return func() {
		x = x + 1
		fmt.Printf("foo2 val = %d\n", x)
	}
}

func main() {
	var i = 1
	bar := foo(i)
	bar()
	bar()
	bar()
	//闭包变量传递
	//func foo1(x *int) func()
	var x = 100
	bar1 := foo1(&x)
	bar1()    //foo1 val = 101
	x = 10000 //因为是引用传递x的作用范围也不局限于foo1，所以当x在其他地方发生变化的时候bar1内部的x也会变化，变得不再封闭
	bar1()    //foo1 val = 10001
	bar1()    //foo1 val = 10002
	bar2 := foo1(&x)
	bar2() //foo1 val = 10003
	bar2() //foo1 val = 10004
	bar2() //foo1 val = 10005
	fmt.Println("-----------------")
	//func foo2(x int) func()
	var x2 = 100
	bar3 := foo2(x2)
	bar3()     //foo1 val = 101
	x2 = 10000 //这里修改是不会影响bar3()闭包内的x值，因为是值传递
	bar3()     //foo1 val = 102
	bar3()     //foo1 val = 103
	bar4 := foo2(x2)
	bar4() //foo1 val = 10001
	bar4() //foo1 val = 10002
	bar4() //foo1 val = 10003

	//闭包的延迟绑定
	foo0()() // 猜猜我会输出什么？ foo0 val = 11

	//goroutine的延迟绑定，本质上也是闭包的延迟绑定问题
	//参考 https://zhuanlan.zhihu.com/p/92634505
	foo5()
	/*
		foo5 val = 5
		foo5 val = 5
		foo5 val = 5
		foo5 val = 5
	*/
	time.Sleep(time.Second * 2)
}

func foo0() func() {
	x := 1
	f := func() {
		fmt.Printf("foo0 val = %d\n", x)
	}
	x = 11
	return f
}

func foo5() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
		go func() {
			fmt.Printf("foo5 val = %v\n", val)
		}()
		//time.Sleep(time.Second) sleep后就是正常的了，所以实际上是因为循环太快了
	}
}
