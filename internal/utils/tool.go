package utils

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"syscall"
	"time"
	"unicode/utf16"
	"unsafe"
)

// StrPtr string到uintptr.
func StrPtr(s string) uintptr {
	p, _ := syscall.BytePtrFromString(s)
	return uintptr(unsafe.Pointer(p))
}

// PtrStr uintptr到string.
func PtrStr(p uintptr) string {
	return syscall.UTF16ToString(*(*[]uint16)(unsafe.Pointer(&p)))
}

// Uint16SliceDataPtr uint16[0]指针到uintptr.
func Uint16SliceDataPtr(p *[]uint16) uintptr {
	if len(*p) == 0 {
		return uintptr(0)
	}
	return uintptr(unsafe.Pointer(&(*p)[0]))
}

// FontInfoName 转换[32]uint16到string
func FontInfoName(str [32]uint16) string {
	return syscall.UTF16ToString(str[0:])
}

// BoolPtr bool到uintptr.
func BoolPtr(b bool) uintptr {
	if b {
		return uintptr(1)
	}
	return uintptr(0)
}

// BoolPtrPtr bool指针到uintptr.
func BoolPtrPtr(b *bool) uintptr {
	return uintptr(unsafe.Pointer(&b))
}

// Float32Ptr float32到uintptr.
func Float32Ptr(f float32) uintptr {
	return uintptr(*(*uint32)(unsafe.Pointer(&f)))
}

// PtrToFloat32 uintptr到float32.
func PtrToFloat32(p uintptr) float32 {
	u := uint32(p)
	return *(*float32)(unsafe.Pointer(&u))
}

// ByteSliceDataPtr byte[0]指针到uintptr.
func ByteSliceDataPtr(b *[]byte) uintptr {
	if len(*b) == 0 {
		return uintptr(0)
	}
	return uintptr(unsafe.Pointer(&(*b)[0]))
}

func PtrStr2(r uintptr) string {
	p := (*uint16)(unsafe.Pointer(r))
	if p == nil {
		return ""
	}
	n, end, add := 0, unsafe.Pointer(p), unsafe.Sizeof(*p)
	for *(*uint16)(end) != 0 {
		end = unsafe.Add(end, add)
		n++
	}
	return string(utf16.Decode(unsafe.Slice(p, n)))
}

// IntPtr int到uintptr
func IntPtr(i int) uintptr {
	return uintptr(i)
}

// TimeFormat 转换为Y-m-d H:i:s时间
func TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

// GbkToUtf8 GBK 转 UTF-8
func GbkToUtf8(s string) (string, error) {
	b := []byte(s)
	reader := transform.NewReader(bytes.NewReader(b), simplifiedchinese.GBK.NewDecoder())
	d, err := io.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(d), nil
}
