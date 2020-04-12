package main

import (
	"fmt"
	ge "github.com/urakozz/go-emoji"
)

func main() {
	aa := "😭asdf a阿斯顿发款式大方😊😄"
	xig := "🍉"
	p := ge.NewEmojiParser()
	s := p.ReplaceAll([]byte(aa), []byte(xig))
	fmt.Println(string(s))
	return
}
