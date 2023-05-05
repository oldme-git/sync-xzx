package utils

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"strconv"
	"strings"
	"time"
)

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

// IsNumber 判断一个字符串是不是数字
func IsNumber(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

// SliceTrim 去除切片多余的空格
func SliceTrim(m []string) {
	for k, v := range m {
		m[k] = strings.TrimSpace(v)
	}
}
