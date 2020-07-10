package main

import (
	"fmt"
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
}
