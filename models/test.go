package models

import (
	"fmt"
)

type Test struct {
	Response string `json:"resp"`
}

func TestFunction() Test {
	fmt.Println("asdf")
	var tv Test
	tv.Response = "Hello World"
	return tv
}
