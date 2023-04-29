package save

import (
	"bufio"
	"changeme/utils"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type AccountInfo struct {
	Name    string
	Sex     int
	DepCode string
	CardId  string
	StudNo  string
}

// Read 获取文件中的用户信息
func Read() (data []AccountInfo) {
	// 读取目录下的文件
	s, err := filepath.Glob("../RecvTemp/*")
	if err != nil {
		panic(err)
	}
	file, err := os.Open(s[0])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 读取文件载入到切片中
	scanner := bufio.NewScanner(file)
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
			data = append(data, AccountInfo{
				Name:    name,
				Sex:     sex,
				DepCode: b[2],
				CardId:  b[3],
				StudNo:  b[4],
			})
		}
	}
	if len(data) > 0 {
		data = data[1:]
	}
	return
}
