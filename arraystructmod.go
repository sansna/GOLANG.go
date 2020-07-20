package main

import (
	"fmt"
	goproto "github.com/gogo/protobuf/proto"
)

type A struct {
	// this can be modified
	AA *int64
	// this cannot be modified
	AAA int64
}
type B struct {
	// this can be modified
	BB []*A
	// this cannot be modified, unless inner ptr fields
	BBB []A
}

func main() {
	b := B{
		BB: []*A{
			&A{
				AA:  goproto.Int64(10),
				AAA: 10,
			},
			&A{
				AA:  goproto.Int64(20),
				AAA: 20,
			},
		},
		BBB: []A{
			A{
				AA:  goproto.Int64(10),
				AAA: 10,
			},
			A{
				AA:  goproto.Int64(20),
				AAA: 20,
			},
		},
	}

	for _, x := range b.BB {
		*x.AA *= 2
		x.AAA *= 2
	}
	for _, x := range b.BBB {
		*x.AA *= 2
		x.AAA *= 2
	}

	fmt.Println(b.BB[0].AAA, b.BB[1].AAA)
	fmt.Println(b.BBB[0].AAA, b.BBB[1].AAA)

	fmt.Println(*b.BB[0].AA, *b.BB[1].AA)
	fmt.Println(*b.BBB[0].AA, *b.BBB[1].AA)
}
