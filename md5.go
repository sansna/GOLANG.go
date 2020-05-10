package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	data := "你好呀我是水电费水电费"
	hasher := md5.New()
	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))
	hasher.Write([]byte(data))
	fmt.Println(hasher.BlockSize())
	hasher.Write([]byte(data))
	fmt.Println(hasher.BlockSize())
	fmt.Println(hex.EncodeToString(hasher.Sum(nil)))

	newhash := md5.New()
	hasher.Write([]byte(""))
	fmt.Println(hex.EncodeToString(newhash.Sum(nil)))
	hasher.Write([]byte(""))
	fmt.Println(hex.EncodeToString(newhash.Sum(nil)))
}
