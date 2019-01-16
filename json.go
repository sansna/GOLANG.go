package main
import (
	"encoding/json"
	"fmt"
)

type a struct {
	a int `json:"a"`
	b int `json:"b"`
}

func main() {
	s := ""
	c := a{}
	err := json.Unmarshal([]byte(s), &c)
	if err != nil {
		fmt.Println(err)
	}
	
}
