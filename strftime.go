package main

import (
	"fmt"
	"time"

	sf "github.com/jehiah/go-strftime"
)

func main() {
	now := time.Now()
	fmt.Println(now.Unix())
	fmt.Println(sf.Format("%Y年%m月%d日 %H时%M分%S秒", now))
	return
}
