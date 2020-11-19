package main

import (
	"fmt"

	//"git.ixiaochuan.cn/pp_server/live/proto/mproto"
	"git.ixiaochuan.cn/pp_server/live/service/common/cfgst"
	"git.ixiaochuan.cn/pp_server/live/service/common/cmodel"
	//"git.ixiaochuan.cn/pp_server/live/service/account/service/statsvc"
)

//"""
//CREATE TABLE `json` (
//	`id` bigint(20) NOT NULL AUTO_INCREMENT,
//	`doc` json DEFAULT NULL,
//	`et` bigint(20) NOT NULL DEFAULT '0',
//	PRIMARY KEY (`id`)
//)
//"""

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

type UserStatus struct {
	Mid    int64  `json:"mid"`
	TypeId string `json:"type_id,omitempty"`
}

func (UserStatus) TableName() string {
	return "user_status"
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

	sql := fmt.Sprintf("select * from %s where doc->\"$.mid\" = ?", doc.TableName())
	has, err = session.SQL(sql, 82).Get(&doc)
	fmt.Println("ok", has, err, doc)

	cnt, _ := session.SQL("select count(*) from json").Count()
	fmt.Println(cnt)

	fmt.Println(session.Table("user_status").Where("type_id = ?", "kladsjfi").Sum(JsonSt{}, "mid"))

	us := UserStatus{}
	session.Where("mid = 12").Get(&us)
	fmt.Println(us.Mid, us.TypeId)
}
