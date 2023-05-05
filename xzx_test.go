package main

import (
	"fmt"
	"sync-xzx/internal/data"
	"testing"
)

func TestXzxInit3(t *testing.T) {
	var s = "123"
	fmt.Println(s[1:2])
}

func TestXzxA(t *testing.T) {
	d := data.New()
	res, err := d.SaveData()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
