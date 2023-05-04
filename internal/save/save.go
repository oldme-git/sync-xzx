package save

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type AccountInfo struct {
	Name    string
	Sex     int
	DepCode string
	CardId  string
	StudNo  string
}

// Save 保存数据到数据库中
func Save(data []AccountInfo) error {
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
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.10.47:3306)/oldme")
	if err != nil {
		return nil, err
	}
	return db, err
}
