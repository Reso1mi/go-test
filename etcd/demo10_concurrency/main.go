package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"log"
	"time"
)

func main() {
	var (
		config clientv3.Config
		client *clientv3.Client
	)
	config = clientv3.Config{
		Endpoints:   []string{"120.79.182.28:2379"},
		DialTimeout: 5 * time.Second,
	}
	client, e := clientv3.New(config)
	if e != nil {
		log.Fatal(e.Error())
	}
	go scheduler(client, "A")
	go scheduler(client, "B")
	for {
		time.Sleep(time.Second)
	}
}

func scheduler(client *clientv3.Client, name string) {
	//创建租约
	response, e := client.Grant(context.TODO(), 10)
	if e != nil {
		log.Fatal(e.Error())
	}
	//创建一个Session自动给租约续租，将租约和Session绑定
	session, e := concurrency.NewSession(client, concurrency.WithLease(response.ID))
	if e != nil {
		log.Fatal(e.Error())
	}
	defer session.Close()
	//创建分布式锁
	mutex := concurrency.NewMutex(session, "/lock")
	//mutex := concurrency.NewMutex(session, name) 这样就是两把锁了
	//创建5s过期的context
	timeout, _ := context.WithTimeout(context.Background(), 5*time.Second)
	e = mutex.Lock(timeout)
	if e != nil {
		//超时了没有获取到锁
		fmt.Println(name, " give up lock")
		return
	}
	fmt.Println(name, "acquire lock")
	//10s的任务，让另一个ctx超时取消
	time.Sleep(10 * time.Second)
	defer func() {
		if e == nil {
			mutex.Unlock(context.TODO())
			fmt.Println(name, "release lock")
		}
	}()
}
