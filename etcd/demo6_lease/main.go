package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main() {
	var (
		config         clientv3.Config
		client         *clientv3.Client
		err            error
		lease          clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseID        clientv3.LeaseID
		putResponse    *clientv3.PutResponse
		getResp        *clientv3.GetResponse
		keepResp       *clientv3.LeaseKeepAliveResponse
		keepRespChan   <-chan *clientv3.LeaseKeepAliveResponse
		kv             clientv3.KV
	)

	config = clientv3.Config{
		Endpoints:   []string{"120.79.182.28:2379"},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}
	//申请一个租约
	lease = clientv3.NewLease(client)
	if leaseGrantResp, err = lease.Grant(context.TODO(), 10); err != nil {
		fmt.Println(err)
		return
	}

	leaseID = leaseGrantResp.ID
	//每秒自动续租
	if keepRespChan, err = lease.KeepAlive(context.TODO(), leaseID); err != nil {
		fmt.Println(err)
		return
	}
	//消费chan
	go func() {
		for {
			select {
			case keepResp = <-keepRespChan:
				if keepResp == nil {
					fmt.Println("租约失效")
					goto END
				} else {
					fmt.Println("收到自动续租应答", keepResp.ID)
				}
			}
		}
	END:
	}()

	kv = clientv3.NewKV(client)
	if putResponse, err = kv.Put(context.TODO(), "/cron/lock/job1", "", clientv3.WithLease(leaseID)); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("写入成功", putResponse.Header.Revision)

	for {
		if getResp, err = kv.Get(context.TODO(), "/cron/lock/job1"); err != nil {
			fmt.Println(err)
			return
		}
		if getResp.Count == 0 {
			fmt.Println("kv过期")
			return
		}
		fmt.Println("还没过期", getResp.Kvs)
		time.Sleep(time.Second * 2)
	}
}
