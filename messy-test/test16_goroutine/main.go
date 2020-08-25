package main

import (
	"fmt"
	_ "net/http/pprof"
	"os"
	"time"
)

func main() {
	for i := 0; i < 30; i++ {
		//执行线程的数量是不定的，根据需要创建
		//writefile.go 启动30个协程不断地写文件。底层会创建30+的线程
		//go WriteFile(i)
		//协程越多，执行线程未必越多，取决于于协程是否忙碌，忙碌的协程越多，执行线程就越多
		//下面这个在linux上创建的线程就没有那么多，大概就2，3条因为都在休眠
		go sleep()
	}
	for {
		time.Sleep(time.Second * 1)
	}
}

func sleep() {
	for {
		time.Sleep(time.Second)
	}
}

func WriteFile(num int) {
	file := fmt.Sprintf("%d.txt", num)
	fp, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR, 0666)
	if nil != err {
		fmt.Printf("openFile failed, err:%s\n", err.Error())
		return
	}
	data := "Hello"
	for {
		fp.Write([]byte(data))
	}
}
