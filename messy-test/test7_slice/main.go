package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

//零切片，空切片，nil切片
func main() {
	var s1 []int
	var s2 = []int{}
	var s3 = make([]int, 0)
	// new 函数返回是指针类型，所以需要使用 * 号来解引用
	var s4 = *new([]int)

	fmt.Println(len(s1), len(s2), len(s3), len(s4)) //0 0 0 0
	fmt.Println(cap(s1), cap(s2), cap(s3), cap(s4)) //0 0 0 0
	fmt.Println(s1, s2, s3, s4)                     //[] [] [] []

	var a1 = *(*[3]int)(unsafe.Pointer(&s1))
	var a2 = *(*[3]int)(unsafe.Pointer(&s2))
	var a3 = *(*[3]int)(unsafe.Pointer(&s3))
	var a4 = *(*[3]int)(unsafe.Pointer(&s4))
	fmt.Println(a1) //[0 0 0] nil切片
	fmt.Println(a2) //[5918120 0 0] 空切片
	fmt.Println(a3) //[5918120 0 0] 空切片
	fmt.Println(a4) //[0 0 0] nil切片

	var s5 = make([]struct{ x, y, z int }, 0)
	var a5 = *(*[3]int)(unsafe.Pointer(&s5))
	// base address for all 0-byte allocations
	//var zerobase uintptr
	fmt.Println(a5) //[824634539336 0 0]

	var se = make([]int, 0, 9)
	fmt.Println(((*reflect.SliceHeader)(unsafe.Pointer(&se))).Len)
	fmt.Println(((*reflect.SliceHeader)(unsafe.Pointer(&se))).Data)
	//append为什么一定要接收返回值呢？
	//写Java写多了，这里一直没转过弯来，最开始有一种说法，说是底层数组要扩容，所以必须要接收
	//这种说法我也不能接受，我一直认为传递的slice进去，就可以在函数内部修改slice的属性，即使扩容了我也可以直接修改底层数组的指针
	//但是我上面的想法其实犯了一个致命的错误，go都是值传递，而slice本身也只是一个普通的struct，包含了Len,Cap,Data几个属性
	//所以它的传递也是通过值传递的，传递的不过是一份拷贝罢了，根本就无法修改原本的slice
	//同时先不说是否扩容的事情，即使不扩容，append势必会使得Len属性变化，所以即使不扩容我们仍然需要接收新的slice
	//那么问题来了，没有扩容，那么Data地址肯定没有变化，但是为什么读取原始的slices无法获取新的值呢？
	//很明显就是因为Len的原因啦，Len限制了切片访问底层数组的长度，所以始终理解一定，切片只是对数组的封装，实际上也是个普通的struct

	//其实还有一种情况，需要使用切片指针，就是把slice做函数参数传递，并在函数中append，函数传递本身就是值传递，传递的是slice的拷贝
	//加上上面的结论，如果不使用指针，即使在函数中接收了append的返回值，对原本的slice也没有任何影响
	//具体代码见 http://imlgw.top/2019/11/06/leetcode-er-cha-shu/#257-%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E6%89%80%E6%9C%89%E8%B7%AF%E5%BE%84
	k := append(se, 1, 2, 3, 4, 5)
	fmt.Println(((*reflect.SliceHeader)(unsafe.Pointer(&se))).Len)
	fmt.Println(((*reflect.SliceHeader)(unsafe.Pointer(&se))).Data)
	fmt.Println(k)
}
