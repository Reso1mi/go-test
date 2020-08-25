package main

import (
	"fmt"
	"time"
)

type User struct {
	name string
	age  int
}

func main() {
	timer := time.NewTimer(time.Second * 5)
	fmt.Println(time.Now())
	select {
	case <-timer.C:
		//default: 有default就直接过去了，否则就一直等，等哪一个先来

	}
	fmt.Println(time.Now())

	/*
		var u = User{name: "imlgw", age: 11}
		time.AfterFunc(time.Duration(3)*time.Second, func(user User) func() {
			return func() {
				fmt.Println(user) //{imlgw 11}
			}
		}(u))
		u.name = "resolmi" //立即执行时复制了一份（结构体复制），所以没有影响
		time.Sleep(5 * time.Second)
	*/
	/*
		var u = &User{name: "imlgw", age: 11}
		time.AfterFunc(time.Duration(3)*time.Second, func(user *User) func() {
			return func() {
				fmt.Println(user) //{resolmi 11}
			}
		}(u))
		u.name = "resolmi" //立即执行时复制了指针，指向同一份数据，所以当外部修改时，内部也会有影响
		//u = User{}
		time.Sleep(5 * time.Second)
	*/
	var u = &User{name: "imlgw", age: 11}
	time.AfterFunc(time.Duration(2)*time.Second, func(user *User) func() {
		return func() {
			fmt.Println(user) //&{imlgw 11}
		}
	}(u))
	u = nil //仅仅改变指针的指向，对闭包内部的指针是没有影响的（立即执行时复制了一份引用，所以没有影响）
	time.Sleep(5 * time.Second)

	//和上面立即执行的写法对比
	var u2 = &User{name: "imlgw", age: 11}
	time.AfterFunc(time.Duration(2)*time.Second, func() {
		fmt.Println(u2) //<nil>
	})
	//u2 = nil //改变外部指针的指向，对闭包内部的指针是有影响的（没有复制u2）
	time.Sleep(5 * time.Second)
}
