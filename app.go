package main

import (
	"context"
	"encoding/json"
	"fmt"
	wruntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"io"
	"os"
	"sync-xzx/internal/conf"
	"sync-xzx/internal/data"
	"sync-xzx/internal/logger"
	"sync-xzx/internal/xzx"
	"time"
)

// App struct
type App struct {
	ctx context.Context
	// 配置文件名
	confName string
	// 关闭自动同步的通道使用
	autoSyncChan chan bool
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		confName: "./conf.json",
	}
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

// DbConnTest 数据库连接测试
func (a *App) DbConnTest() string {
	db, err := data.Open()
	if err != nil {
		return "连接失败：" + err.Error()
	}
	if err := db.Ping(); err != nil {
		return "连接失败：" + err.Error()
	}
	return "数据库连接成功"
}

// SaveConf 保存配置
func (a *App) SaveConf(s string) string {
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		return err.Error()
	}

	c := conf.NewConf(a.confName)

	// 写入数据
	err = c.Write(m)
	if err != nil {
		return err.Error()
	}
	return "保存配置成功"
}

// ReadConf 读取配置
func (a *App) ReadConf() string {
	file, err := os.Open(a.confName)
	if err != nil {
		return err.Error()
	}
	defer file.Close()

	s, err := io.ReadAll(file)
	if err != nil {
		return err.Error()
	}

	return string(s)
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
			// 保存数据
			d := data.New(context.Background())
			go func() {
				for ready := range d.Progress() {
					// 向客户端主动推动当前进度
					wruntime.EventsEmit(a.ctx, "progress", ready)
				}
				return
			}()
			res, err := d.SaveData()
			if err != nil {
				return "保存数据失败：" + err.Error()
			}
			return fmt.Sprintf("获取到%d条数据，保存成功%d条数据", num, res)
		}
		return fmt.Sprintf("获取失败%d", inqAllAccRes)
	} else {
		return fmt.Sprintf("init3失败%d", initRes)
	}
}

// StartAutoSync 开启自动同步
func (a *App) StartAutoSync() {
	a.autoSyncChan = make(chan bool, 1)
	// 读取同步时间
	go func() {
		//wruntime.EventsEmit(a.ctx, "autoSyncStatus", "启动自动同步，同步时间"+ext["syncTime"].(string)+"分钟")
		for range time.Tick(1 * time.Second) {
			select {
			case <-a.autoSyncChan:
				// 关闭自动同步
				close(a.autoSyncChan)
				wruntime.EventsEmit(a.ctx, "autoSyncStatus", "关闭自动同步")
				return
			default:
				// 正在同步
				fmt.Println(1)
			}
		}
	}()
}

// CloseAutoSync 开启自动同步
func (a *App) CloseAutoSync() {
	a.autoSyncChan <- true
}
