package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

type CronJob struct {
	expr     *cronexpr.Expression
	nextTime time.Time //expr.Next(now)
}

func main() {
	var (
		cronJob       *CronJob
		expr          *cronexpr.Expression
		now           time.Time
		scheduleTable map[string]*CronJob
	)
	scheduleTable = make(map[string]*CronJob)
	now = time.Now()
	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	scheduleTable["job1"] = cronJob

	now = time.Now()
	expr = cronexpr.MustParse("*/5 * * * * * *")
	cronJob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}
	scheduleTable["job2"] = cronJob

	//启动一个调度协程
	go func() {
		for {
			now = time.Now()
			for name, job := range scheduleTable {
				if job.nextTime.Before(now) || job.nextTime.Equal(now) {
					go func(name string) {
						fmt.Println("执行", name)
					}(name)
					job.nextTime = job.expr.Next(now)
					fmt.Println(name, "下次执行时间：", job.nextTime)
				}
			}
			select {
			case <-time.NewTimer(100 * time.Millisecond).C: //将在100ms可读，返回

			}
		}
	}()
	time.Sleep(time.Second * 100)
}
