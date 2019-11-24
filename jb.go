package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"

	jb "github.com/yanyiwu/gojieba"
)

func main() {
	x := jb.NewJieba()
	defer x.Free()
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		s := s.Text()
		words := x.Cut(s, true)
		fmt.Println(strings.Join(words, "\n"))
	}
}
