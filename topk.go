package main

import (
	"fmt"
	"unsafe"

	topk "github.com/dgryski/go-topk"
)

func main() {
	s := topk.New(100)
	s.Insert("4", 1)
	s.Insert("1", 1)
	s.Insert("1", 8)
	s.Insert("1", 1)
	s.Insert("1", 1)
	s.Insert("2", 1)
	s.Insert("9", 9)
	s.Insert("4", 1)
	fmt.Println(s.Keys()[0].Key)
	res, _ := s.GobEncode()
	fmt.Printf("encoded: %v\n", unsafe.Sizeof(res))
	news := &topk.Stream{}
	news.GobDecode(res)
	fmt.Println("decoded:", news.Keys()[1].Key)
	//
	res = []byte("")
	news = &topk.Stream{}
	news.GobDecode(res)
	fmt.Println("empty decode: ", news.Keys())
}
