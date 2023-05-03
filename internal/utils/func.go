package utils

import (
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
)

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

func Log(s string) {

}
