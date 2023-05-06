package main

import (
	"context"
	"fmt"
	"runtime"
	"sync-xzx/internal/conf"
	"sync-xzx/internal/data"
	"testing"
	"time"
)

// 测试保存数据，需要在该目录下存有RecvTemp和ControlFile
func TestSave(t *testing.T) {
	fmt.Println("初始协程数量", runtime.NumGoroutine())
	ctx := context.Background()
	d := data.New(ctx)
	go func() {
		for ready := range d.Progress() {
			fmt.Println("读到数据", ready)
		}
		return
	}()
	res, err := d.SaveData()

	if err != nil {
		panic(err)
	}
	fmt.Println("最终数据条数", res)
	time.Sleep(2 * time.Second)
	fmt.Println("结束协程数量", runtime.NumGoroutine())
}

// 测试保存数据的性能
func BenchmarkSave(b *testing.B) {
	ctx := context.Background()
	d := data.New(ctx)
	go func() {
		for _ = range d.Progress() {
		}
		return
	}()
	d.SaveData()
}

func TestA(t *testing.T) {
	c := conf.NewConf("conf.json")
	ext, err := c.GetConf("ext")
	if err != nil {
		panic(err)
	}
	fmt.Println(ext)
}
