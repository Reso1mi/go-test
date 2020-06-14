package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"time"
)

func main() {
	var (
		config             clientv3.Config
		client             *clientv3.Client
		err                error
		kv                 clientv3.KV
		watcher            clientv3.Watcher
		getResp            *clientv3.GetResponse
		watchStartRevision int64
		watchRespChan      clientv3.WatchChan
		watchResp          clientv3.WatchResponse
		event              *clientv3.Event
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

	go func() {
		for {
			kv.Put(context.TODO(), "/cron/jobs/job1", "{...job1...}")
			kv.Delete(context.TODO(), "/cron/jobs/job1")
			time.Sleep(time.Second * 1)
		}
	}()

	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job1"); err != nil {
		fmt.Println(err)
		return
	}
	//key存在
	if len(getResp.Kvs) != 0 {
		fmt.Println("当前值:", string(getResp.Kvs[0].Value))
	}
	//当前etcd集群事务ID,单调递增
	watchStartRevision = getResp.Header.Revision + 1
	fmt.Println("从当前版本向后监听：", watchStartRevision)
	//创建监听器
	watcher = clientv3.NewWatcher(client)
	//创建一个context
	ctx, cancelFunc := context.WithCancel(context.TODO())
	time.AfterFunc(5*time.Second, func() {
		cancelFunc() //5s后撤销监听
	})
	watchRespChan = watcher.Watch(ctx, "/cron/jobs/job1", clientv3.WithRev(watchStartRevision))
	//监听kv变化事件
	for watchResp = range watchRespChan {
		for _, event = range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("修改为：", string(event.Kv.Value), "Revision:", event.Kv.CreateRevision, event.Kv.ModRevision)
			case mvccpb.DELETE:
				fmt.Println("删除了", "Revision:", event.Kv.ModRevision)
			}
		}
	}

}
