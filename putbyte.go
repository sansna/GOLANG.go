package main
import (
	"encoding/binary"
	"fmt"
)

func main() {
	a := int64(1919865)
	key := make([]byte,16)
	binary.BigEndian.PutUint64(key, uint64(a))
	fmt.Println(key)
}
