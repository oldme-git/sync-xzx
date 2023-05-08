// 一个简单的定时任务，用于实现自动同步

package cron

import (
	"context"
	"time"
)

type Cron struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func AddCron(ctx context.Context, t time.Duration, f func(ctx context.Context)) *Cron {
	ctx, cancel := context.WithCancel(ctx)
	ticker := time.NewTicker(t)

	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				f(ctx)
			}
		}
	}()

	return &Cron{
		ctx:    ctx,
		cancel: cancel,
	}
}

// IsRunning 定时任务是否还在运行
func (c *Cron) IsRunning() bool {
	return c.ctx.Err() == nil
}

func (c *Cron) Exit() {
	c.cancel()
}
