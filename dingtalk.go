package main

import (
	"git.ixiaochuan.cn/xclib/common/utils"
	"strconv"
)

func main() {
	text := "好奇o"
	text += "\n手机号:" + ""
	for _, iid := range []int64{1, 2, 3, 4} {
		text += "\nhttp://file.iupvideo.net/view/png/id/" + strconv.FormatInt(iid, 10)
	}
	utils.DingDingSendText("https://oapi.dingtalk.com/robot/send?access_token=", "反馈:"+text, []string{})
}
