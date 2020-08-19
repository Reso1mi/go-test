package main

import (
	"context"
	"fmt"
	"time"
)

/*
	Context 有两个主要的功能：
	1.通知子协程退出（正常退出，超时退出等）；
	2.传递必要的参数。
*/

type Options struct{ Interval time.Duration }

func reqTask(ctx context.Context, name string) {
	for {
		select {
		//在子协程中，使用 select 调用 <-ctx.Done() 判断是否需要退出
		//Done会返回一个chan 里面是空的struct{}
		case <-ctx.Done():
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			//获取上下文中携带的信息
			v := ctx.Value("options").(*Options)
			time.Sleep(v.Interval * 1 * time.Second)
		}
	}
}

func main() {
	//context.Background(): 创建根 Context，通常在 main 函数、初始化和测试代码中创建，作为顶层 Context。
	//context.WithCancel(parent) 创建可取消的子Context，同时返回函数 cancel。
	ctx, cancel := context.WithCancel(context.Background())
	//也可以创建带定时器的context
	//ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	//context.WithDeadline() 也可以设置截至时间

	//基于上面的ctx再创建一个指vCtx，这个可以携带值
	vCtx := context.WithValue(ctx, "options", &Options{1})
	go reqTask(vCtx, "worker1")
	//也可以同时控制多个协程
	go reqTask(vCtx, "worker2")
	time.Sleep(3 * time.Second)
	//主协程中，调用 cancel() 函数通知子协程退出。
	cancel()
	time.Sleep(1 * time.Second)
	/*  result:
	worker1 send request
	worker2 send request
	worker1 send request
	worker2 send request
	worker1 send request
	worker2 send request
	stop worker2
	stop worker1
	*/

	/*
		#.1
		通过下面4个函数创建context.Background()根节点的子孙context
		这四个函数的第一个参数都是父context，返回一个Context类型的值
		func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
		func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
		func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
		func WithValue(parent Context, key interface{}, val interface{}) Context

		#.2
		context上下文数据的存储就像一个树，每个结点只存储一个key/value对。
		WithValue()保存一个key/value对，它将父context嵌入到新的子context，并在节点中保存了key/value数据。
		Value()查询key对应的value数据，会从当前context中查询，如果查不到，会递归查询父context中的数据。
		值得注意的是，context中的上下文数据并不是全局的，它只查询本节点及父节点们的数据，不能查询兄弟节点的数据。
	*/
}
