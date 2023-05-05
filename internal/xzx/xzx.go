package xzx

import (
	"os"
	"sync-xzx/internal/conf"
	"sync-xzx/internal/logger"
	"sync-xzx/internal/utils"
	"syscall"
	"unsafe"
)

// dll初始化
var xzxApi = syscall.NewLazyDLL("AIO_API.dll")

// Init3 新中新初始化
func Init3() (int, error) {
	var (
		init3 = xzxApi.NewProc("TA_Init3")
		c     = conf.NewConf("conf.json")
	)
	xzx, err := c.GetDb("xzx")
	if err != nil {
		return 0, err
	}

	var (
		host       = utils.StrToPtrChar(xzx["ip"].(string))
		port       = utils.StrToInt(xzx["port"].(string))
		sysCode    = utils.StrToInt(xzx["sysCode"].(string))
		terminalNo = utils.StrToInt(xzx["terminal"].(string))
		// 是否脱机
		proxyOffline = utils.BoolToPtrBool(false)
		// 交易流水
		maxJnl = utils.IntToPrtInt(999999)
	)

	// 调用接口
	res, _, _ := init3.Call(host, port, sysCode, terminalNo, proxyOffline, maxJnl)
	return int(res), nil
}

// InqAllAcc 读取所有账户信息
func InqAllAcc() (int, int, error) {
	// 删除之前的数据
	if err := os.RemoveAll("./RecvTemp"); err != nil {
		logger.NewLogger().Log(err.Error())
	}
	var (
		c        = conf.NewConf("conf.json")
		xzx, err = c.GetDb("xzx")
	)
	if err != nil {
		return 0, 0, err
	}
	var (
		taInqAllAcc = xzxApi.NewProc("TA_InqAllAcc")
		recNum      int
		fileName    = utils.StrToPtrChar("DSQ")
		timeOut     = utils.StrToInt(xzx["timeout"].(string))
	)

	res, _, _ := taInqAllAcc.Call(uintptr(unsafe.Pointer(&recNum)), fileName, timeOut)
	return int(res), recNum, nil
}

// ExtractConFileDep 读取部门信息
func ExtractConFileDep() bool {
	// 删除之前的数据
	var (
		taIExtractConFile = xzxApi.NewProc("TA_ExtractConFile")
		readrec           = utils.IntToInt(0)
		fn                = utils.StrToPtrChar("dep.txt")
	)

	res, _, _ := taIExtractConFile.Call(readrec, fn)
	return uintptr(res) == 1
}

// ExtractConFilePid 读取身份
func ExtractConFilePid() bool {
	// 删除之前的数据
	var (
		taIExtractConFile = xzxApi.NewProc("TA_ExtractConFile")
		readrec           = utils.IntToInt(1)
		fn                = utils.StrToPtrChar("pid.txt")
	)

	res, _, _ := taIExtractConFile.Call(readrec, fn)
	return uintptr(res) == 1
}
