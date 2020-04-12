package main

import (
	"fmt"
	"github.com/ajg/form"
	"strings"
)

func main() {
	//str := "AccountID=kcs600001&&Race=asdf"
	str := "AccountID=kcs600001"
	type A struct {
		Aid  string `form:"AccountID,omitempty"`
		Race string `form:"Race,omitempty"`
	}
	a := &A{}
	d := form.NewDecoder(strings.NewReader(str))
	err := d.Decode(a)
	if err != nil {
		fmt.Print("err", err)
	} else {
		fmt.Println(a.Aid)
		fmt.Println(a)
	}
}
