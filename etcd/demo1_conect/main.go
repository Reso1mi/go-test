package main

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
		err    error
	)
	config = clientv3.Config{
		Endpoints:   []string{"120.79.182.28:2379"},
		DialTimeout: 5 * time.Second,
	}
	//这里其实测试不出来连接上没有，这里创建好client之后client内部会用协程进行重连
	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}
	client = client
}
