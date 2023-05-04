package main

import (
	"changeme/internal/conf"
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// 保存到数据库
func (a *App) save() string {
	err := saveData()
	if err != nil {
		return err.Error()
	}
	return ""
}

// 同步新中新数据
func (a *App) sync() string {
	res := xzxInit()
	if res == 0 {
		res2, num := inqAllAcc()
		if res2 == 0 {
			return fmt.Sprintf("获取到%d条数据\n", num)
		}
	}
	return "err"
}

// SaveConf 保存配置
func (a *App) SaveConf(s string) string {
	c := conf.NewConf("./conf.json")

	// 写入数据
	err := c.Write(s)
	if err != nil {
		return err.Error()
	}
	return "保存成功"
}
