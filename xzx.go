package main

import (
	"changeme/internal/utils"
	"unsafe"
)

// 初始化
func xzxInit() int {
	var (
		init3 = xzxApi.NewProc("TA_Init3")
		host  = utils.StrPtr(ip)
		// 是否脱机
		prop         = false
		proxyOffline = utils.BoolPtrPtr(&prop)
		// 交易流水
		maxJnlUint = make([]uint16, 999999)
		maxJnl     = utils.Uint16SliceDataPtr(&maxJnlUint)
	)

	// 调用接口
	res, _, _ := init3.Call(host, uintptr(port), uintptr(sysCode), uintptr(terminalNo), proxyOffline, maxJnl)
	return int(res)
}

func inqAllAcc() (int, int) {
	var (
		taInqAllAcc = xzxApi.NewProc("TA_InqAllAcc")
		recNum      int
		fileName    = utils.StrPtr("DSQ")
		timeOut     = utils.IntPtr(10)
	)

	res, _, _ := taInqAllAcc.Call(uintptr(unsafe.Pointer(&recNum)), fileName, timeOut)
	return int(res), recNum
}
