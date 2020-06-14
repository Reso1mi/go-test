package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type cmdResult struct {
	output []byte
	err    error
}

func main() {
	var (
		ctx        context.Context
		cancelFunc context.CancelFunc
		cmd        *exec.Cmd
		resultChan chan *cmdResult
		result     *cmdResult
	)
	//结果队列
	resultChan = make(chan *cmdResult, 1000)
	ctx, cancelFunc = context.WithCancel(context.TODO())
	go func() {
		var (
			output []byte
			err    error
		)
		//执行任务捕获输出
		cmd = exec.CommandContext(ctx, "bash", "-c", "sleep 2s;echo \"Hello World\"")
		output, err = cmd.CombinedOutput()
		//通过chan传递给main协程
		resultChan <- &cmdResult{
			output: output,
			err:    err,
		}
	}()
	time.Sleep(1 * time.Second)
	//取消上下文
	cancelFunc()
	result = <-resultChan
	fmt.Println(string(result.output), result.err)
}
