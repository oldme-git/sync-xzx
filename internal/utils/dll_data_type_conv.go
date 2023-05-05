// 调用dll的所有数据类型转换为uintptr类型

package utils

import (
	"strconv"
	"syscall"
	"unsafe"
)

// StrToPtrChar string to *char
func StrToPtrChar(s string) uintptr {
	p, _ := syscall.BytePtrFromString(s)
	return uintptr(unsafe.Pointer(p))
}

// StrToPtrInt string to *int
func StrToPtrInt(s string) uintptr {
	i, _ := strconv.Atoi(s)
	return uintptr(unsafe.Pointer(&i))
}

// StrToInt string to int
func StrToInt(s string) uintptr {
	i, _ := strconv.Atoi(s)
	return uintptr(i)
}

// IntToInt int to int
func IntToInt(i int) uintptr {
	return uintptr(i)
}

// IntToPrtInt int to *int
func IntToPrtInt(i int) uintptr {
	return uintptr(unsafe.Pointer(&i))
}

// BoolToBool bool to bool
func BoolToBool(b bool) uintptr {
	if b {
		return uintptr(1)
	}
	return uintptr(0)
}

// BoolToPtrBool bool to *bool
func BoolToPtrBool(b bool) uintptr {
	return uintptr(unsafe.Pointer(&b))
}
