package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string `json:"name"`
	Body string `json:"body"`
	Time int64  `json:"time"`
}

func main() {
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, _ := json.Marshal(m)
	fmt.Println(string(b)) //{"name":"Alice","body":"Hello","time":1294706395881547000}
	var msg1 **Message
	var msg2 = &Message{}
	//很神奇，传msg就不行，传&msg就可以
	//大致看了下源码，首先如果传msg1其实就相当于传了一个nil，在反射阶段就直接报错了
	//而传入指针的地址会通过反射创建一个值 见：decode.go/indirect 488行
	json.Unmarshal(b, &msg1)
	fmt.Println(*msg1)
	json.Unmarshal(b, &msg2)
	fmt.Println(msg2)
}
