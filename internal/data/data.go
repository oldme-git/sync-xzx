package data

import (
	"bufio"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"path/filepath"
	"strings"
	"sync-xzx/internal/utils"
	"time"
)

type Data struct {
	row
	pidList map[string]string
	depList map[string]string
}

type row struct {
	readerNo        string
	readerName      string
	dpCode          string
	dpName          string
	readerType1     string
	readerType1Code string
	readerType2     string
	readerType2Code string
	readerType3     string
	readerType3Code string
	readerDept1     string
	readerDept1Code string
	readerDept2     string
	readerDept2Code string
	readerDept3     string
	readerDept3Code string
	registerYear    string
	cardNo          string
	cardStatus      string
	expireDate      int
	updateTime      int64
	oldCardNo       string
}

func New() *Data {
	var data Data
	data.injectPid()
	data.injectDep()
	return &data
}

// SaveData 读取并保存数据到库中
func (d *Data) SaveData() (res int, err error) {
	// 读取目录下的文件
	s, err := filepath.Glob("./RecvTemp/*")
	if err != nil {
		return
	}
	if len(s) == 0 {
		return 0, errors.New("RecvTemp dir not find file")
	}
	file, err := os.Open(s[0])
	if err != nil {
		return
	}
	defer file.Close()

	// 读取文件载入到切片中
	var (
		scanner   = bufio.NewScanner(file)
		dataSlice []row
	)
	for scanner.Scan() {
		var (
			line    = strings.Split(scanner.Text(), "|")
			lineRow row
		)
		if len(line) > 1 {
			utils.SliceTrim(line)
			name, _ := utils.GbkToUtf8(line[0])
			// 部门标识抛弃前三位
			line[2] = line[2][3:]
			// 数据载入
			lineRow = row{
				readerNo:        line[5],
				readerName:      name,
				readerType1:     d.pidList[line[7]],
				readerType1Code: line[7],
				cardNo:          line[3],
				updateTime:      time.Now().Unix(),
			}
			// 第一级部门
			if len(line[2]) >= 6 {
				lineRow.readerDept1 = d.depList[line[2][:6]]
				lineRow.readerDept1Code = line[2][:6]

				// 第二级部门
				if len(line[2]) >= 9 {
					lineRow.readerDept2 = d.depList[line[2][:9]]
					lineRow.readerDept2Code = line[2][:9]

					// 第三级部门
					if len(line[2]) >= 12 {
						lineRow.readerDept3 = d.depList[line[2]]
						lineRow.readerDept3Code = line[2]
					}
				}
			}

			dataSlice = append(dataSlice, lineRow)
		}
	}
	if len(dataSlice) > 0 {
		dataSlice = dataSlice[1:]
	}
	res, err = save(dataSlice)
	if err != nil {
		return
	}
	return
}

// 注入身份
func (d *Data) injectPid() {
	file, err := os.Open("./ControlFile/pid.txt")
	if err != nil {
		d.pidList = nil
	}
	defer file.Close()

	// 读取文件载入到切片中
	var (
		scanner = bufio.NewScanner(file)
		m       = make(map[string]string)
	)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		if len(line) > 4 {
			utils.SliceTrim(line)
			line[3], _ = utils.GbkToUtf8(line[3])
			m[line[0]] = line[3]
		}
	}
	d.pidList = m
}

// 注入部门
func (d *Data) injectDep() {
	file, err := os.Open("./ControlFile/dep.txt")
	if err != nil {
		d.depList = nil
	}
	defer file.Close()

	// 读取文件载入到切片中
	var (
		scanner = bufio.NewScanner(file)
		m       = make(map[string]string)
	)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")
		if len(line) > 2 {
			utils.SliceTrim(line)
			if !utils.IsNumber(line[1]) {
				line[1], _ = utils.GbkToUtf8(line[1])
			}
			m[line[0]] = line[1]
		}
	}
	d.depList = m
}
