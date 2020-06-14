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
		kv             clientv3.KV
		lease          clientv3.Lease
		leaseGrantResp *clientv3.LeaseGrantResponse
		leaseID        clientv3.LeaseID
		keepRespChan   <-chan *clientv3.LeaseKeepAliveResponse
		keepResp       *clientv3.LeaseKeepAliveResponse
		ctx            context.Context
		cancelFUnc     context.CancelFunc
		txn            clientv3.Txn
		txnResp        *clientv3.TxnResponse
	)

	config = clientv3.Config{
		Endpoints:   []string{"120.79.182.28:2379"},
		DialTimeout: 5 * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}
	//lease实现锁自动过期,避免宕机死锁
	//op操作
	//txn事务 if else then

	//##1. 上锁 （创建租约，自动续租，拿租约去抢占key）
	//申请一个租约
	lease = clientv3.NewLease(client)
	//创建context,用于取消自动续租
	ctx, cancelFUnc = context.WithCancel(context.TODO())
	if leaseGrantResp, err = lease.Grant(ctx, 5); err != nil {
		fmt.Println(err)
		return
	}
	leaseID = leaseGrantResp.ID
	//确保函数退出后自动续租会停止，KeepAlive被取消，chan会返回空
	defer cancelFUnc()
	//立即释放租约，关联的kv也会被删除
	defer lease.Revoke(context.TODO(), leaseID)
	//每秒自动续租
	if keepRespChan, err = lease.KeepAlive(context.TODO(), leaseID); err != nil {
		fmt.Println(err)
		return
	}
	//启动协程消费chan
	go func() {
		for {
			select {
			case keepResp = <-keepRespChan:
				if keepResp == nil {
					fmt.Println("租约失效")
					goto END //退出协程
				} else {
					fmt.Println("收到自动续租应答", keepResp.ID)
				}
			}
		}
	END:
	}()
	//抢占key
	kv = clientv3.NewKV(client)
	//创建事务
	txn = kv.Txn(context.TODO())
	//定义事务
	txn.If(clientv3.Compare(clientv3.CreateRevision("/cron/lock/job9"), "=", 0)).
		Then(clientv3.OpPut("/cron/lock/job9", "{lock}", clientv3.WithLease(leaseID))).
		Else(clientv3.OpGet("/cron/lock/job9"))
	if txnResp, err = txn.Commit(); err != nil {
		fmt.Println(err)
		return //可以直接return 有defer机制
	}
	if !txnResp.Succeeded {
		fmt.Println("锁被占用", string(txnResp.Responses[0].GetResponseRange().Kvs[0].Value))
		return
	}

	//##2. 处理业务
	// 进入分布式锁，很安全
	fmt.Println("处理任务....")
	time.Sleep(10 * time.Second)
	//##3. 释放锁（取消自动续租，释放租约）
}
