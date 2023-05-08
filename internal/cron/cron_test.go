package cron

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestCron(t *testing.T) {
	fmt.Println("初始协程数量", runtime.NumGoroutine())
	var (
		tick = 2 * time.Second
		ctx  = context.Background()
	)
	cron := AddCron(ctx, tick, func(ctx context.Context) {
		fmt.Println("cron执行了")
	})

	time.Sleep(5 * time.Second)
	fmt.Println("关闭cron")
	cron.Exit()

	time.Sleep(10 * time.Second)
	fmt.Println("结束协程数量", runtime.NumGoroutine())
}
