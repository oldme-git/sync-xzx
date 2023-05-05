package conf

import (
	"encoding/json"
	"io"
	"os"
)

type Conf struct {
	filename string
}

// NewConf 创建一个新的 Conf 实例
func NewConf(filename string) *Conf {
	return &Conf{filename}
}

// 写入数据到文件
func (c *Conf) Write(data interface{}) error {
	file, err := os.Create(c.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

// 从文件中读取数据
func (c *Conf) Read(data interface{}) error {
	file, err := os.Open(c.filename)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonData, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonData, data)
	if err != nil {
		return err
	}

	return nil
}

// GetDb 获取conf
// t是配置类型，当前有db,xzx,ext
func (c *Conf) GetDb(t string) (m map[string]interface{}, err error) {
	// 读取数据
	var data map[string]interface{}
	if err := c.Read(&data); err != nil {
		return nil, err
	}
	m = data[t].(map[string]interface{})
	return
}
