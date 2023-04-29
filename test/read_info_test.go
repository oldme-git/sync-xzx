package test

import (
	"changeme/save"
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	data := save.Read()
	fmt.Println(data)
}
