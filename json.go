package main

import (
	"encoding/json"
	"fmt"
)

type a struct {
	A int `json:"a"`
	B int `json:"b"`
	D int `json:"d"`
}

type b struct {
	D int `json:"d"`
	E int `json:"e"`
	A a   `json:"z"`
}

func main() {
	s := ""
	c := a{}
	err := json.Unmarshal([]byte(s), &c)
	if err != nil {
		fmt.Println(err)
	}

	aa := a{
		A: 1,
		B: 2,
		D: 3,
	}
	buf, err := json.Marshal(aa)
	if err != nil {
		fmt.Println("err marshal.")
		return
	}
	bb := b{}
	err = json.Unmarshal(buf, &bb)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bb.A)
}
