package main
import (
	"time"
	"fmt"
)

func main() {
	fmt.Println(time.Now().AddDate(0,0,10).String())
}
