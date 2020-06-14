package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	var (
		config  clientv3.Config
		client  *clientv3.Client
		err     error
		kv      clientv3.KV
		delResp *clientv3.DeleteResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"120.79.182.28:2379"},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}
	//读写etcd的键值对
	kv = clientv3.NewKV(client)
	if delResp, err = kv.Delete(context.TODO(), "/cron/jobs/", clientv3.WithPrefix()); err != nil {
		fmt.Println(err)
	} else {
		if len(delResp.PrevKvs) != 0 {
			for _, kvPair := range delResp.PrevKvs {
				fmt.Println("delete", string(kvPair.Value))
			}
		}
	}
}
