package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Unix(-100000,0).Year())
}
