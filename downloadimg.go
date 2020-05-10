package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	resp, _ := http.Get("http://zyimg101.oss-cn-shanghai.aliyuncs.com/zyimg/51/c4/49f9-8081-11ea-8171-00163e007bff?x-oss-process=image/watermark,image_enlpbWcvZTYvMjgvNTA1ZS05OWQ0LTExZTktODAwYS0wMDE2M2UwODI3OTU_eC1vc3MtcHJvY2Vzcz1pbWFnZS9yZXNpemUsbF8xNTAvZm9ybWF0LHBuZy9jaXJjbGUscl8xNTA=,g_center")
	buf, _ := ioutil.ReadAll(resp.Body)
	f, _ := os.Create("a.png")
	f.Write(buf)
	f.Close()
}
