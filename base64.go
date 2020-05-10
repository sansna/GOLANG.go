package main

import (
	"bytes"
	"os"
	"fmt"
	"encoding/base64"
)

func main() {
	input := []byte("zyimg/e6/28/505e-99d4-11e9-800a-00163e082795?x-oss-process=image/resize,l_150/format,png/circle,r_150")
	encoder := base64.NewEncoder(base64.URLEncoding, os.Stdout)
	encoder.Write(input)
	encoder.Close()
	fmt.Println()

	var buf bytes.Buffer
	input = []byte("zyimg/96/1f/2590-80b4-11ea-8045-00163e082f12?x-oss-process=image/resize,l_140/format,png/circle,r_140")
	enc := base64.NewEncoder(base64.URLEncoding, &buf)
	enc.Write(input)
	enc.Close()
	fmt.Println(buf.String())
}
