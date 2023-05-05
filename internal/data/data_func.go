package data

import (
	"database/sql"
	"sync-xzx/internal/logger"
)

// 保存数据到数据库中
func save(data []row) (res int, err error) {
	// 打开数据库
	db, err := open()
	if err != nil {
		return
	}
	// 重置数据表
	_, err = db.Exec(`DROP TABLE IF EXISTS sz_reader`)
	if err != nil {
		return
	}
	_, err = db.Exec(`CREATE TABLE sz_reader (
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
	stmt, err := db.Prepare("INSERT INTO sz_reader(reader_no, reader_name, dp_code, dp_name, reader_type1, reader_type1_code, reader_type2, reader_type2_code, reader_type3, reader_type3_code, reader_dept1, reader_dept1_code, reader_dept2, reader_dept2_code, reader_dept3, reader_dept3_code, register_year, card_no, card_status, expire_date, update_time, old_card_no) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return
	}
	defer stmt.Close()

	for _, v := range data {
		_, err := stmt.Exec(
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
			res++
		}
	}

	return
}

// 打开数据库
func open() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.10.47:3306)/oldme")
	if err != nil {
		return nil, err
	}
	return db, err
}
