package main

import (
	//"reflect"
	"bytes"
	"encoding/gob"
	"fmt"
)

type A struct {
	A int64
	B int32
}

type Msg struct {
	App   string
	Mid   string
	Type  string
	Stype string
	Data  map[string]interface{}
}

func main() {
	//init data
	a := A{
		A: 10,
		B: 20,
	}
	str := "helloworld"
	// new types should be registerd first hand
	gob.Register(A{})
	data := map[string]interface{}{
		"asdf":     "Asdfsadl",
		"int":      12,
		"a":        a,
		"slicestr": []string{"123", "sdkf", "zxj"},
		"lpstr":    &str,
		"xcv":      nil,
	}
	msg := Msg{
		App:  "hanabi",
		Mid:  "123",
		Type: "chat",
		Data: data,
	}

	// init dec/enc method
	var encmsg bytes.Buffer
	enc := gob.NewEncoder(&encmsg)
	dec := gob.NewDecoder(&encmsg)

	// encoding
	err := enc.Encode(msg)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	strmsg := encmsg.String()
	//fmt.Println(strmsg)

	// decoding
	encmsg = *bytes.NewBufferString(strmsg)
	var z Msg
	// ensure encmsg is not empty
	if encmsg.Len() > 0 {
		err = dec.Decode(&z)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	}
	// type of z is Msg ?!
	//fmt.Println(reflect.TypeOf(z))
	fmt.Println(z.Data["xcv"])
}
