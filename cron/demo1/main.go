package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	var (
		expr     *cronexpr.Expression
		err      error
		now      time.Time
		nextTime time.Time
	)
	if expr, err = cronexpr.Parse("* * * * *"); err != nil {
		fmt.Println(err)
	}

	//每隔5分钟执行一次,支持的粒度比linux更精细
	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
	}
	now = time.Now()
	nextTime = expr.Next(now)
	fmt.Println(nextTime)
	//next-now 时间差
	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("被调度了", nextTime)
	})
	time.Sleep(30 * time.Second)
}
