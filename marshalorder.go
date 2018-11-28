package main

import "fmt"
import proto "git.ixiaochuan.cn/xclib/adapter/proto/post/record"
import "git.ixiaochuan.cn/service/post/record/controllers"

// json.marshal keeps order of slice
func main() {
	controllers.Init()
	ids := []int64{1, 2, 3, 99, 63, 10, -1, 0, 7, 9}
	data := proto.GetRecordSquareIdsData{Ids: ids}
	c := &controllers.HttpController{}
	fmt.Println(c.EncodeReply(data))
}
