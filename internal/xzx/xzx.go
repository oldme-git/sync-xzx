package xzx

import (
	"errors"
	"fmt"
	"os"
	"sync-xzx/internal/conf"
	"sync-xzx/internal/logger"
	"sync-xzx/internal/utils"
	"syscall"
	"unsafe"
)

type Xzx struct {
	xzxApi *syscall.LazyDLL
	// xzx的配置
	conf map[string]interface{}
}

// New 创建一个xzx连接
func New() (*Xzx, error) {
	var (
		xzx = &Xzx{}
		err error
	)

	// dll初始化
	xzx.xzxApi = syscall.NewLazyDLL("AIO_API.dll")
	if err := xzx.xzxApi.Load(); err != nil {
		utils.ErrLog()
		return nil, err
	}

	// 读取配置
	xzx.conf, err = conf.NewConf("conf.json").GetConf("xzx")
	if err != nil {
		return nil, err
	}

	// init3初始化
	err = xzx.init3()
	if err != nil {
		return nil, err
	}

	return xzx, nil
}

// Init3 新中新初始化
func (x *Xzx) init3() (err error) {
	var (
		init3      = x.xzxApi.NewProc("TA_Init3")
		host       = utils.StrToPtrChar(x.conf["ip"].(string))
		port       = utils.StrToInt(x.conf["port"].(string))
		sysCode    = utils.StrToInt(x.conf["sysCode"].(string))
		terminalNo = utils.StrToInt(x.conf["terminal"].(string))
		// 是否脱机
		proxyOffline = utils.BoolToPtrBool(false)
		// 交易流水
		maxJnl = utils.IntToPrtInt(999999)
	)

	// 调用接口
	res, _, _ := init3.Call(host, port, sysCode, terminalNo, proxyOffline, maxJnl)
	if int(res) != 0 {
		return errors.New(fmt.Sprintf("调用TA_Init3失败了,错误码%d", int(res)))
	}
	return
}

// InqAllAcc 读取所有账户信息
func (x *Xzx) InqAllAcc() (int, error) {
	// 删除之前的数据
	if err := os.RemoveAll("./RecvTemp"); err != nil {
		logger.NewLogger().Log(err.Error())
	}

	var (
		taInqAllAcc = x.xzxApi.NewProc("TA_InqAllAcc")
		recNum      int
		fileName    = utils.StrToPtrChar("DSQ")
		timeOut     = utils.StrToInt(x.conf["timeout"].(string))
	)

	res, _, _ := taInqAllAcc.Call(uintptr(unsafe.Pointer(&recNum)), fileName, timeOut)
	if int(res) != 0 {
		return 0, errors.New(fmt.Sprintf("调用TA_InqAllAcc失败了,错误码%d", int(res)))
	}
	return recNum, nil
}

// ExtractConFileDep 读取部门信息
func (x *Xzx) ExtractConFileDep() error {
	// 删除之前的数据
	var (
		taIExtractConFile = x.xzxApi.NewProc("TA_ExtractConFile")
		readrec           = utils.IntToInt(0)
		fn                = utils.StrToPtrChar("dep.txt")
	)

	res, _, _ := taIExtractConFile.Call(readrec, fn)
	if uintptr(res) == 1 {
		return nil
	} else {
		return errors.New("调用TA_ExtractConFile读取部门失败")
	}
}

// ExtractConFilePid 读取身份
func (x *Xzx) ExtractConFilePid() error {
	// 删除之前的数据
	var (
		taIExtractConFile = x.xzxApi.NewProc("TA_ExtractConFile")
		readrec           = utils.IntToInt(1)
		fn                = utils.StrToPtrChar("pid.txt")
	)

	res, _, _ := taIExtractConFile.Call(readrec, fn)
	if uintptr(res) == 1 {
		return nil
	} else {
		return errors.New("调用TA_ExtractConFile读取部门失败")
	}
}
