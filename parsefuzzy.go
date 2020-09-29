package main

import (
	//"encoding/json"
	"fmt"
	"github.com/json-iterator/go"

	"github.com/json-iterator/go/extra"
)

func main() {
	extra.RegisterFuzzyDecoders()

	type A struct {
		AA string `json:"aa" msgpack:"aa"`
	}
	str := "{\"aa\":10}"
	a := &A{}

	err := jsoniter.UnmarshalFromString(str, a)
	fmt.Println(err)
	fmt.Println(a)

	//a.AA = 1
	//byt, err := json.Marshal(a)
	//fmt.Println(string(byt), err)
}
