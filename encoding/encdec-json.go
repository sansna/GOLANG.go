package main

import (
	//"strconv"
	"encoding/json"
	"fmt"
)

type Msg struct {
	App   string
	Mid   string
	Type  string
	Stype string
	Data  map[string]interface{}
}

func main() {
	data := map[string]interface{}{
		"das": "sdfas\nsdfs",
		"int": 12,
	}
	a := Msg{
		App:  "hanabi",
		Mid:  "123",
		Type: "post",
		Data: data,
	}
	byt, _ := json.Marshal(a)
	result := Msg{}
	json.Unmarshal(byt, &result)
	fmt.Printf("%v", result)
}
