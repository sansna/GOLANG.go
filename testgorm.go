package main

import (
	"git.ixiaochuan.cn/pp_server/live/service/common/cfgst"
	"git.ixiaochuan.cn/pp_server/live/service/common/cmodel"
)

type Account struct {
	Id int64 `json:"id,omitempty" gorm:"column:id"`
	Et int64 `json:"et,omitempty" gorm:"column:et"`
}

func (Account) TableName() string {
	return "json"
}

type AXX struct {
	A int
	B int
	C int
	D int
}

func (AXX) TableName() string {
	return "axx"
}

func main() {
	cmodel.InitMysqlGORM(cfgst.MySql{
		Hostport: "127.0.0.1:3306",
		User:     "root",
		Password: "zz",
		DbName:   "account",
	}, true)

	session := cmodel.DBNew()
	// this does set et = 100
	session.Model(Account{}).Where("id = ?", 2).Select("et").Updates(&Account{
		Et: 100,
	})
	// this does not set et = 0, gorm ignores blank fields
	session.Model(Account{}).Where("id = ?", 2).Select("et").Updates(&Account{
		Et: 0,
	})

	if !session.HasTable(AXX{}) {
		cmodel.CreateTables(session, []cmodel.TableDef{
			cmodel.TableDef{
				Table: AXX{},
				Indexes: [][]string{
					[]string{"a", "b", "c"},
					[]string{"b", "c"},
				},
				UniqIndexes: [][]string{
					[]string{"c", "d"},
				},
			},
		})
	}

}
