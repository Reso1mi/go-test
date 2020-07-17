package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	//Go中的字符串是只读的，相当于一个只读的切片
	var str = "abcdefghijklmn"
	//str[0] = 'a' //Cannot assign to str[0]
	st := str[1:] //bcdefghijklmn
	fmt.Println(str, st)

	s := "hello, world"
	s1 := "hello, world"
	s2 := "hello, world"[7:]
	//3者指向相同的内存地址，也就是说这里的字面量是公用的，类似Java的字符串池，Java中的String同样也是不可变的
	fmt.Printf("%d \n", (*reflect.StringHeader)(unsafe.Pointer(&s)).Data)  // 5043540
	fmt.Printf("%d \n", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Data) // 5043540
	fmt.Printf("%d \n", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Data) // 5043540+7 = 5043547

	//当我们使用for range迭代字符串时，
	//每次迭代Go都会用UTF-8解码出一个rune类型(int32)的字符，且索引为当前rune的起始位置(以字节为最下单位)。
	for index, char := range "你好" {
		fmt.Printf("start at %d, Unicode = %U, char = %c\n", index, char, char)
		//start at 0, Unicode = U+4F60, char = 你
		//start at 3, Unicode = U+597D, char = 好
	}

	//随机访问一个字符串的时候，返回的是单个字节byte(uint8)，和上面for range不一样
	ss := "你好"
	//s[1] = '½', hex = bd, Unicode = U+00BD
	fmt.Printf("s[%d] = %q, hex = %x, Unicode = %U", 1, ss[1], ss[1], ss[1])
}
