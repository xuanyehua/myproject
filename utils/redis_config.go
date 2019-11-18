package utils

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/config"
	_ "github.com/gomodule/redigo/redis"
)

func Get_Redis()  (adapter cache.Cache, err error){
	appconf,err := config.NewConfig("ini","./conf/app.conf")
	if err != nil{
		fmt.Println("数据库配置异常",err)
		return
	}
	fmt.Println("aaaaaaaa",appconf.String("redis_addr")+":"+appconf.String("redis_prost"))
	//redis_config_map := map[string]string{
	//	"conn":"127.0.0.1:6379",
	//	"dbNum" : "1",
	//}
	redis_config_map := map[string]string{
		"conn":appconf.String("redis_addr")+":"+appconf.String("redis_prost"),
		"dbNum" : appconf.String("redis_db"),
	}
	redis_config ,_:=json.Marshal(redis_config_map)
	bm ,err := cache.NewCache("redis",string(redis_config))
	return bm ,err
}
