package conf

import (
	"fmt"
	"testing"
)

func TestWrite(t *testing.T) {
	conf := NewConf("data.json")

	// 写入数据
	err := conf.Write(map[string]string{"ip": "127.0.0.1", "port": "8500"})
	if err != nil {
		panic(err)
	}
}

func TestRead(t *testing.T) {
	conf := NewConf("data.json")

	// 读取数据
	var data map[string]interface{}
	err := conf.Read(&data)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}

func TestGetConf(t *testing.T) {
	conf := NewConf("data.json")
	sub, _ := conf.GetConf("db")
	fmt.Println(sub["dbName"])
}

func TestSetConf(t *testing.T) {
	conf := NewConf("data.json")
	err := conf.SetConf("db", "dbName", "traceint")
	if err != nil {
		panic(err)
	}
}
