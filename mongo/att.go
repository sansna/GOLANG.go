package main

import (
	"fmt"
	"strings"

	"github.com/wangbin/jiebago/analyse"
	"gopkg.in/mgo.v2/bson"

	"git.ixiaochuan.cn/service/base"
	"git.ixiaochuan.cn/xclib/common/mongodb"
	"git.ixiaochuan.cn/xclib/common/mongodb/cfgstruct"
	"git.ixiaochuan.cn/xclib/common/xcproto"
	rds_cfg "git.ixiaochuan.cn/xclib/common/xcredis/cfgstruct"
)

var (
	Acntmgo   *mongodb.MultiplexClient
	Attmgo    *mongodb.MultiplexClient
	Yonmgo    *mongodb.MultiplexClient
	Postmgo   *mongodb.MultiplexClient
	Recordmgo *mongodb.MultiplexClient
	Matchmgo  *mongodb.MultiplexClient
	Revmgo    *mongodb.MultiplexClient
	Chatmgo   *mongodb.MultiplexClient
	Mediamgo  *mongodb.MultiplexClient
	Oprdwmgo  *mongodb.MultiplexClient
	Oprmgo    *mongodb.MultiplexClient
)

var (
	Acntcache   *base.Cache
	Recordcache *base.Cache
	Chatcache   *base.Cache
	Revcache    *base.Cache
	Attcache    *base.Cache
	Yoncache    *base.Cache
	Matchcache  *base.Cache
	Oprcache    *base.Cache
)

var err error

func main() {
	acntmgo, err := mongodb.InitMultiplexV2(cfgstruct.MongodbSt{
		Hostport: "dds-uf65797e0d7307141.mongodb.rds.aliyuncs.com:3717",
		UserName: "hanabi",
		Password: "Dei7cai4uYookoo",
		Poolsize: 8,
	}, "account")
	if err != nil {
		fmt.Println("err: %v", err)
		return
	}
	attmgo, err := mongodb.InitMultiplexV2(cfgstruct.MongodbSt{
		Hostport: "dds-uf6bd416d125f1541.mongodb.rds.aliyuncs.com:3717",
		UserName: "hanabi",
		Password: "Dei7cai4uYookoo",
		Poolsize: 8,
	}, "attention")
	if err != nil {
		fmt.Println("err: %v", err)
		return
	}
	acntcache, err := base.NewCacheV2(rds_cfg.RedisSt{
		Hostport: "r-uf6934ae6d85c604.redis.rds.aliyuncs.com:6379",
		Poolsize: 8,
		Timeout:  30,
		Auth:     "crypt-124e9177ac8b466ee3c8c6713cf7a8f06704ed13952238941ba9ecd1db83f840",
	}, "account-redis")

	attcache, err := base.NewCacheV2(rds_cfg.RedisSt{
		Hostport: "r-uf6a32057fc47fd4.redis.rds.aliyuncs.com:6379",
		Poolsize: 8,
		Timeout:  30,
		Auth:     "crypt-124e9177ac8b466ee3c8c6713cf7a8f06704ed13952238941ba9ecd1db83f840",
	}, "attention-redis")

}
