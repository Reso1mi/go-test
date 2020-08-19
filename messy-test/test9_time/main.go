package main

import (
	"fmt"
	"time"
)

func main() {
	var t time.Time
	fmt.Println(t)
	var d time.Duration
	fmt.Println(d)
	var a int
	fmt.Println(a)
	//until类似sub，比sub更好记（shorthand）？
	//等价于time.Now().Add(3*time.Second).Sub(time.Now())
	dur := time.Until(time.Now().Add(3 * time.Second))
	fmt.Println(dur) //3s
}
