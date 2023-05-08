package main

import (
	"context"
	"encoding/json"
	"fmt"
	wruntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"io"
	"os"
	"strconv"
	"sync-xzx/internal/conf"
	"sync-xzx/internal/cron"
	"sync-xzx/internal/data"
	"sync-xzx/internal/logger"
	"sync-xzx/internal/utils"
	"sync-xzx/internal/xzx"
	"time"
)

// App struct
type App struct {
	ctx context.Context
	// 配置文件名
	confName string
	// 自动同步的定时任务
	cron *cron.Cron
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
	_, err := xzx.New()
	if err != nil {
		return err.Error()
	}
	return "连接成功"
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
	x, err := xzx.New()
	if err != nil {
		return err.Error()
	}
	recNum, err := x.InqAllAcc()
	if err != nil {
		return err.Error()
	}
	// 同步部门信息和身份信息
	if err := x.ExtractConFileDep(); err != nil {
		logger.NewLogger().Log(err.Error())
	}
	if err := x.ExtractConFilePid(); err != nil {
		logger.NewLogger().Log(err.Error())
	}
	saveRes := a.Save()
	return fmt.Sprintf("获取到%d条数据，%s", recNum, saveRes)
}

// Save 将文件中的新中新数据直接放入数据库
func (a *App) Save() string {
	// 保存数据
	d := data.New(a.ctx)
	// 向客户端主动推动当前进度
	d.Progress(func(ready int) {
		wruntime.EventsEmit(a.ctx, "progress", ready)
	})
	saveNum, err := d.Save()
	if err != nil {
		return "保存数据失败：" + err.Error()
	}
	return fmt.Sprintf("保存成功%d条数据", saveNum)
}

// StartAutoSync 开启自动同步
func (a *App) StartAutoSync() string {
	// 保存配置
	c := conf.NewConf("conf.json")
	err := c.SetConf("ext", "isAutoSync", "1")
	if err != nil {
		return err.Error()
	}
	ext, err := c.GetConf("ext")
	if err != nil {
		return err.Error()
	}

	// 开启自动同步
	tick, err := strconv.Atoi(ext["syncTime"].(string))
	if err != nil {
		return err.Error()
	}
	if a.cron != nil && a.cron.IsRunning() {
		return "已经有一个定时任务了"
	}

	// 正式开启
	a.cron = cron.AddCron(a.ctx, time.Duration(tick)*time.Minute, func(ctx context.Context) {
		wruntime.EventsEmit(ctx, "autoSyncStatus", "开始同步")
		a.sendSyncTime(time.Duration(tick))
		res := a.Sync()
		wruntime.EventsEmit(ctx, "autoSyncStatus", res)
	})
	a.sendSyncTime(time.Duration(tick))
	return fmt.Sprintf("启动自动同步，间隔%s分钟", ext["syncTime"].(string))
}

// CloseAutoSync 关闭自动同步
func (a *App) CloseAutoSync() string {
	// 保存配置
	c := conf.NewConf("conf.json")
	err := c.SetConf("ext", "isAutoSync", "0")
	if err != nil {
		return err.Error()
	}
	if a.cron != nil && a.cron.IsRunning() {
		a.cron.Exit()
		return "关闭自动同步"
	}
	return "没有需要关闭的同步任务"
}

func (a *App) sendSyncTime(tick time.Duration) {
	// 记录本次同步时间和下一次需要同步的时间
	timeMap := map[string]string{
		"prev": utils.TimeFormat(time.Now()),
		"next": utils.TimeFormat(time.Now().Add(tick * time.Minute)),
	}
	timeJson, _ := json.Marshal(timeMap)
	wruntime.EventsEmit(a.ctx, "autoTime", string(timeJson))
}
