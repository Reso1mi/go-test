package main

import (
	"bytes"
	"context"
	"fmt"
	_ "net/http/pprof"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	/*	go func() {
			err := http.ListenAndServe(":6060", nil)
			if err != nil {
				fmt.Printf("failed to start pprof monitor:%s", err)
			}
		}()
		ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
		defer cancelFn()
		cmd := exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 50; echo hello")
		output, err := cmd.CombinedOutput()
		fmt.Printf("output：【%s】err:【%s】", string(output), err)*/

	ctx, cancelFn := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFn()
	cmd := exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 50; echo hello")
	output, err := RunCmd(ctx, cmd)
	fmt.Printf("output：【%s】err:【%s】", string(output), err)
}

func RunCmd(ctx context.Context, cmd *exec.Cmd) ([]byte, error) {
	var (
		b   bytes.Buffer
		err error
	)
	//开辟新的线程组（Linux平台特有的属性）
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true, //使得Shell进程开辟新的PGID,即Shell进程的PID,它后面创建的所有子进程都属于该进程组
	}
	cmd.Stdout = &b
	cmd.Stderr = &b
	if err = cmd.Start(); err != nil {
		return nil, err
	}
	var finish = make(chan struct{})
	defer close(finish)
	go func() {
		select {
		case <-ctx.Done(): //超时/被cancel 结束
			//kill -(-PGID)杀死整个进程组
			syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
		case <-finish: //正常结束
		}
	}()
	//wait等待goroutine执行完，然后释放FD资源
	//这个时候再kill掉shell进程就不会再等待了，会直接返回
	if err = cmd.Wait(); err != nil {
		return nil, err
	}
	return b.Bytes(), err
}
