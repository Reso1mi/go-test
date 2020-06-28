package main

import "fmt"

//每次 const 出现时，都会让 iota 初始化为0.【自增长】
const (
	x = iota
	_
	y
	z = "zz" //中断iota
	k
	p = iota + 1 //恢复，恢复后从5开始 5+1=6
	m            //7
	n = iota + 2 //9 注意不是8，这里相当于重置重新赋值，所以并不是从8开始，而是从7开始
)

const (
	a = 1
	b
	c
)

// https://www.cnblogs.com/zsy/p/5370052.html
func main() {
	fmt.Println(a, b, c)             //1 1 1
	fmt.Println(x, y, z, k, p, m, n) //0 2 zz zz 6 7 9
}
