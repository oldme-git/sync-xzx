package main

import (
	"bufio"
	"changeme/internal/save"
	"changeme/internal/utils"
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// 读取并保存数据到库中
func saveData() (err error) {
	// 读取目录下的文件
	s, err := filepath.Glob("./RecvTemp/*")
	if err != nil {
		return err
	}
	if len(s) == 0 {
		return errors.New("RecvTemp dir not find file")
	}
	file, err := os.Open(s[0])
	if err != nil {
		return err
	}
	defer file.Close()

	// 读取文件载入到切片中
	var (
		scanner = bufio.NewScanner(file)
		data    []save.AccountInfo
	)
	for scanner.Scan() {
		var (
			line = scanner.Text()
			b    = strings.Split(line, "|")
		)
		if len(b) > 4 {
			var (
				name, _ = utils.GbkToUtf8(b[0])
				sex, _  = strconv.Atoi(b[1])
			)
			data = append(data, save.AccountInfo{
				Name:    name,
				Sex:     sex,
				DepCode: b[2],
				CardId:  b[3],
				StudNo:  b[5],
			})
		}
	}
	if len(data) > 0 {
		data = data[1:]
	}
	err = save.Save(data)
	if err != nil {
		return err
	}
	return
}
