package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()
		if len([]rune(s)) > 20 {
			fmt.Println(s)
		}
	}
	return
}
