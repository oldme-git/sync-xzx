package main

import (
	"context"
	"encoding/json"
	"fmt"
	"sync-xzx/internal/conf"
	"sync-xzx/internal/logger"
	"sync-xzx/internal/xzx"
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

// XzxConnTest 新中新连接测试
func (a *App) XzxConnTest() string {
	res, err := xzx.Init3()
	if err != nil {
		return "连接失败：" + err.Error()
	}
	if res == 0 {
		return "连接成功"
	}
	return fmt.Sprintf("新中新连接失败,返回值%d", res)
}

// SaveConf 保存配置
func (a *App) SaveConf(s string) string {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		return err.Error()
	}

	c := conf.NewConf("./conf.json")

	// 写入数据
	err = c.Write(m)
	if err != nil {
		return err.Error()
	}
	return "保存成功"
}

// Sync 立即同步新中新数据
func (a *App) Sync() string {
	initRes, err := xzx.Init3()
	if err != nil {
		return "init3失败：" + err.Error()
	}
	if initRes == 0 {
		inqAllAccRes, num, err := xzx.InqAllAcc()
		if err != nil {
			return "InqAllAcc失败：" + err.Error()
		}
		if inqAllAccRes == 0 {
			// 同步部门信息和身份信息
			if b := xzx.ExtractConFileDep(); !b {
				logger.NewLogger().Log("同步部门信息失败")
			}
			if b := xzx.ExtractConFilePid(); !b {
				logger.NewLogger().Log("同步身份信息失败")
			}
			return fmt.Sprintf("获取到%d条数据", num)
		}
		return fmt.Sprintf("获取失败%d", inqAllAccRes)
	} else {
		return fmt.Sprintf("init3失败%d", initRes)
	}
}
