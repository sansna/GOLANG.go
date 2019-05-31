package main

import (
	"fmt"
	"time"
)

func main() {
	now := int64(636480000000 / 1000)
	fmt.Println(time.Unix(now, 0).AddDate(28, 0, 0).Unix())
}
