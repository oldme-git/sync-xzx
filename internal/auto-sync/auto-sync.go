package auto_sync

import (
	"fmt"
	"sync-xzx/internal/conf"
)

func Auto() error {
	var (
		c        = conf.NewConf("conf.json")
		ext, err = c.GetConf("db")
	)
	if err != nil {
		return err
	}
	fmt.Println(ext)
	return nil
}
