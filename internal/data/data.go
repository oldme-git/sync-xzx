package data

import (
	"bufio"
	"context"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync-xzx/internal/logger"
	"sync-xzx/internal/utils"
	"time"
)

type Data struct {
	row
	pidList map[string]string
	depList map[string]string
	ctx     context.Context
	// 用来抛出当前已经保存数据的进度
	progressChan chan int
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

func New(ctx context.Context) *Data {
	var data Data
	data.injectPid()
	data.injectDep()
	data.ctx = ctx
	return &data
}

// SaveData 解析出数据并保存数据到库中
func (d *Data) SaveData() (res int, err error) {
	// 启用进度通道，对外通知当前进度
	d.progressChan = make(chan int, 1)

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
			// 数据载入
			lineRow = row{
				readerNo:        line[5],
				readerName:      name,
				readerType1:     d.pidList[line[7]],
				readerType1Code: line[7],
				cardNo:          line[3],
				updateTime:      time.Now().Unix(),
				cardStatus:      "0",
			}
			// 部门处理
			d.depHandel(line, &lineRow)
			// 卡状态处理
			d.pidHandel(line, &lineRow)

			dataSlice = append(dataSlice, lineRow)
		}
	}
	if len(dataSlice) > 0 {
		dataSlice = dataSlice[1:]
	}

	res, err = d.save(dataSlice)
	if err != nil {
		return
	}
	return
}

// 保存数据到数据库中
func (d *Data) save(data []row) (res int, err error) {
	// 打开数据库
	db, err := Open()
	defer db.Close()

	if err != nil {
		return
	}
	// 重置数据表
	_, err = db.ExecContext(d.ctx, `DROP TABLE IF EXISTS sz_reader`)
	if err != nil {
		return
	}
	_, err = db.ExecContext(d.ctx, `CREATE TABLE sz_reader (
		id int(10) unsigned NOT NULL AUTO_INCREMENT,
		reader_no varchar(100) NOT NULL,
		reader_name varchar(255) DEFAULT NULL,
		dp_code varchar(100) DEFAULT NULL,
		dp_name varchar(100) DEFAULT NULL,
		reader_type1 varchar(100) DEFAULT NULL,
		reader_type1_code varchar(100) DEFAULT NULL,
		reader_type2 varchar(100) DEFAULT NULL,
		reader_type2_code varchar(100) DEFAULT NULL,
		reader_type3 varchar(100) DEFAULT NULL,
		reader_type3_code varchar(100) DEFAULT NULL,
		reader_dept1 varchar(100) DEFAULT '',
		reader_dept1_code varchar(100) DEFAULT NULL,
		reader_dept2 varchar(100) DEFAULT NULL,
		reader_dept2_code varchar(100) DEFAULT NULL,
		reader_dept3 varchar(255) DEFAULT NULL,
		reader_dept3_code varchar(100) DEFAULT NULL,
		register_year varchar(100) DEFAULT NULL,
		card_no varchar(50) DEFAULT '',
		card_status varchar(100) DEFAULT NULL,
		expire_date bigint(16) unsigned DEFAULT NULL,
		update_time int(10) unsigned DEFAULT NULL,
		old_card_no varchar(50) DEFAULT '',
		PRIMARY KEY (id) USING BTREE,
		KEY reader_type (reader_type1) USING BTREE,
		KEY reader_no (reader_no) USING BTREE
	) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='读者信息表(部门和身份信息做冗余)'`)
	if err != nil {
		return
	}

	// 保存数据
	stmt, err := db.PrepareContext(d.ctx, "INSERT INTO sz_reader(reader_no, reader_name, dp_code, dp_name, reader_type1, reader_type1_code, reader_type2, reader_type2_code, reader_type3, reader_type3_code, reader_dept1, reader_dept1_code, reader_dept2, reader_dept2_code, reader_dept3, reader_dept3_code, register_year, card_no, card_status, expire_date, update_time, old_card_no) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return
	}
	defer stmt.Close()

	// 分割切片，开启协程单条插入数据
	var (
		goNum = 10
		wg    = sync.WaitGroup{}
		// res锁
		rw              sync.RWMutex
		insertData      = chunkSlice(data, goNum)
		ctxCurr, cancel = context.WithCancel(d.ctx)
	)

	wg.Add(goNum)
	for _, chuck := range insertData {
		go func(r []row) {
			for _, v := range r {
				_, err = stmt.ExecContext(
					ctxCurr,
					v.readerNo,
					v.readerName,
					v.dpCode,
					v.dpName,
					v.readerType1,
					v.readerType1Code,
					v.readerType2,
					v.readerType2Code,
					v.readerType3,
					v.readerType3Code,
					v.readerDept1,
					v.readerDept1Code,
					v.readerDept2,
					v.readerDept2Code,
					v.readerDept3,
					v.readerDept3Code,
					v.registerYear,
					v.cardNo,
					v.cardStatus,
					v.expireDate,
					v.updateTime,
					v.oldCardNo,
				)
				if err != nil {
					// 记录错误日志
					logger.NewLogger().Log(err.Error())
				} else {
					rw.Lock()
					res++
					rw.Unlock()
				}
			}
			wg.Done()
		}(chuck)
	}

	// 向通道中塞入当前进度
	go func(ctx context.Context) {
		d.progressChan <- res
		for range time.Tick(500 * time.Millisecond) {
			select {
			case <-ctx.Done():
				// 关闭进度通道
				d.progressChan <- res
				close(d.progressChan)
				return
			default:
				d.progressChan <- res
			}
		}
	}(ctxCurr)
	wg.Wait()
	cancel()
	return
}

// Progress 获取进度通道，注意，一旦saveData操作结束了，此通道会关闭
func (d *Data) Progress() chan int {
	return d.progressChan
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

// 部门处理
func (d *Data) depHandel(line []string, lineRow *row) {
	if len(line[2]) < 4 {
		return
	}
	// 部门标识抛弃前三位
	depCode := line[2][3:]
	// 第一级部门
	if len(depCode) >= 6 {
		lineRow.readerDept1 = d.depList[depCode[:6]]
		lineRow.readerDept1Code = depCode[:6]

		// 第二级部门
		if len(depCode) >= 9 {
			lineRow.readerDept2 = d.depList[depCode[:9]]
			lineRow.readerDept2Code = depCode[:9]

			// 第三级部门
			if len(depCode) >= 12 {
				lineRow.readerDept3 = d.depList[depCode]
				lineRow.readerDept3Code = depCode
			}
		}
	}
}

// 卡状态处理
func (d *Data) pidHandel(line []string, lineRow *row) {
	flag := line[12]
	if len(flag) > 3 {
		if flag[1:3] == "00" {
			lineRow.cardStatus = "1"
		}
	} else {
		// 如果flog数据异常，则判断此卡无效
		lineRow.cardStatus = "0"
	}
}
