package save

import (
	"fmt"
	"testing"
)

func TestSave(t *testing.T) {
	err := Save()
	if err != nil {
		panic(err)
	}
}

func TestRead(t *testing.T) {
	data, err := read()
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
