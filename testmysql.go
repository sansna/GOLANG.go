package main

import (
	"fmt"

	//"git.ixiaochuan.cn/pp_server/live/proto/mproto"
	"git.ixiaochuan.cn/pp_server/live/service/common/cfgst"
	"git.ixiaochuan.cn/pp_server/live/service/common/cmodel"
	//"git.ixiaochuan.cn/pp_server/live/service/account/service/statsvc"
)

type JsonItem struct {
	Mid    int64   `json:"mid"`
	Status int64   `json:"status"`
	Uids   []int64 `json:"uids"`
}
type JsonSt struct {
	Id  int64    `json:"id"`
	Doc JsonItem `json:"doc" xorm:"JSON"`
	Et  int64    `json:"et"`
}

func (e JsonSt) TableName() string {
	return "json"
}

func main() {
	cmodel.InitMysql(cfgst.MySql{
		Hostport: "localhost",
		User:     "root",
		Password: "zz",
		DbName:   "account",
	}, true)

	session := cmodel.NewDBSession()
	defer session.Close()

	doc := JsonSt{}
	has, err := session.Where("id=100").Get(&doc)
	fmt.Println(doc, has, err)
	fmt.Println(doc.Doc)
	//stat.AddUserStat(mproto.UserStatusSt{
	//	Mid:     10,
	//	TypeId:  "kladsjfi",
	//	Et:      -100,
	//	BagType: 4,
	//})
	item := JsonSt{
		Id: 101,
		Et: 81728937,
		Doc: JsonItem{
			Mid:    87,
			Status: -101,
			Uids:   []int64{8, 9, 10},
		},
	}
	session.Insert(item)
}
