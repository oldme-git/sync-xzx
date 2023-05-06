package data

import (
	"database/sql"
	"fmt"
	"sync-xzx/internal/conf"
)

// Open 打开数据库
func Open() (*sql.DB, error) {
	c := conf.NewConf("conf.json")
	note, err := c.GetConf("db")
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", note["user"], note["pass"], note["ip"], note["port"], note["dbName"]))
	if err != nil {
		return nil, err
	}
	return db, err
}

// chunkSlice 将一个切片等额的分成chunk块
func chunkSlice(slice []row, chunk int) [][]row {
	var (
		chunks = make([][]row, chunk)
		div    int
	)
	for k, v := range slice {
		div = k % chunk
		chunks[div] = append(chunks[div], v)
	}
	return chunks
}
