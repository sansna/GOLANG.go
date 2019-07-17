package main

import (
	"os"
	//"fmt"
	"encoding/base64"
)

func main() {
	input := []byte("zxcvzxlk")
	encoder := base64.NewEncoder(base64.StdEncoding, os.Stdout)
	encoder.Write(input)
	encoder.Close()
}
