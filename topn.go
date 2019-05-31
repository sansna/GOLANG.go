package main

import (
	"time"
	"math/rand"
	"fmt"

	"./logic"
)

type DATA struct {
	v int
}
type DATAS struct {
	data []DATA
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ds := &DATAS{
		data: make([]DATA, 20),
	}
	for i := 0; i < 20; i++ {
		ds.data[i].v = rand.Intn(300)
	}
	ret := logic.FindN(ds)
	fmt.Println(ret)
	fmt.Println(ds)
}

func (d *DATAS) N() int {
	return 4
}

func (d *DATAS) Len() int {
	return len(d.data)
}

func (d *DATAS) Less(i, j interface{}) bool {
	switch typ1 := i.(type) {
	case DATA:
		switch typ2 := j.(type) {
		case DATA:
			if typ1.v < typ2.v {
				return true
			} else {
				return false
			}
		default:
			return false
		}
	default:
		return false
	}
	return false
}

func (d *DATAS) Get(i int) interface{} {
	return d.data[i]
}
