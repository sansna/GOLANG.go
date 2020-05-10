package main

import (
	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	//var png []byte
	qrcode.WriteFile("https://www.huohuaa.com/download?uk=123asdfk", qrcode.Highest, 750, "a.png")
}
