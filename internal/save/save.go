package save

import (
	"bufio"
	"changeme/internal/utils"
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
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

func Save() error {
	data, err := read()
	if err != nil {
		return err
	}
	// 打开数据库
	db, err := open()
	if err != nil {
		return err
	}
	// 重置数据表
	_, err = db.Exec(`DROP TABLE IF EXISTS xzx_student`)
	if err != nil {
		return err
	}
	_, err = db.Exec(`CREATE TABLE xzx_student (
	  name varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
	  sex tinyint(3) UNSIGNED NULL DEFAULT NULL,
	  dep_code varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
	  card_id varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
	  stud_no varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL
	) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic`)
	if err != nil {
		return err
	}

	// 保存数据
	stmt, err := db.Prepare("INSERT INTO xzx_student(name, sex, dep_code, card_id, stud_no) VALUES(?, ?, ?, ?, ? )")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, v := range data {
		stmt.Exec(v.Name, v.Sex, v.DepCode, v.CardId, v.StudNo)
	}

	return nil
}

// 打开数据库
func open() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:JCPHqknyy8ATR5ME@tcp(192.168.10.47:3306)/oldme")
	if err != nil {
		return nil, err
	}
	return db, err
}

// 获取文件中的用户信息
func read() (data []AccountInfo, err error) {
	// 读取目录下的文件
	s, err := filepath.Glob("../../RecvTemp/*")
	if err != nil {
		return nil, err
	}
	if len(s) == 0 {
		return nil, errors.New("RecvTemp dir not find file")
	}
	file, err := os.Open(s[0])
	if err != nil {
		return nil, err
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
