package main

import (
	"encoding/json"
	"fmt"
	"git.ixiaochuan.cn/service/post/record/server/logic"
)

func AllAutoRecordCheck(fea logic.RecordFeature) bool {
	if fea.Stage == logic.ADM_STAGE_ALL_AUTO && fea.Disp >= logic.MAX_ALL_AUTO_RECORD_DISPLAY {
		return false
	}
	return true
}

func main() {
	//s := "{\"pid\":112599736,\"mid\":8369079,\"tid\":0,\"ct\":1551864043,\"mct\":1477091623,\"fastt\":1551864209,\"baset\":1551867714,\"allt\":1551868885,\"rept\":1551893193,\"liket\":1551868808,\"status\":0,\"forbid\":0,\"stage\":4,\"level\":1,\"age\":20,\"gender\":2,\"v\":3111,\"t\":700,\"u\":6,\"r\":1,\"detail_square\":34,\"detail\":34,\"clen\":34,\"nimg\":2}"
	//ss:="{\"pid\":112600481,\"mid\":153600573,\"tid\":0,\"ct\":1551864293,\"mct\":1550328095,\"fastt\":1551864461,\"baset\":1551867931,\"allt\":1551869137,\"rept\":1551868972,\"liket\":1551894075,\"status\":0,\"forbid\":0,\"stage\":4,\"level\":1,\"age\":18,\"gender\":2,\"v\":10710,\"t\":700,\"u\":28,\"r\":8,\"detail_square\":247,\"detail\":247,\"clen\":22,\"nimg\":1}"
	s:="{\"pid\":112749106,\"mid\":70000607,\"tid\":0,\"ct\":1551925300,\"mct\":1529081582,\"ft\":1551925707,\"fastt\":1551925779,\"rect\":1551925779,\"editrect\":1551925860,\"status\":2,\"forbid\":0,\"stage\":205,\"level\":1,\"city_code\":\"010\",\"age\":1,\"gender\":2,\"v\":0,\"t\":30000\"clen\":413,\"nimg\":1}"
	//fmt.Println(s)
	z := logic.RecordFeature{}
	err := json.Unmarshal([]byte(s), &z)
	if err != nil {
		fmt.Println("err", err)
	}
	if !AllAutoRecordCheck(z) {
		fmt.Println("check pass")
	} else {
		fmt.Println("check not pass")
	}
	fmt.Println(z.Disp)
	fmt.Println(z.Stage)
}
