package utils

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

var (
	Mysql_addr string  //数据库ip
	Mysql_prot  string  // 数据库端口号
	Mysql_db string    // 数据库名称
	Mysql_name string   //账号
	Mysql_psw string    //密码
)


func Inconfig(){
	appconf,err := config.NewConfig("ini","./conf/app.conf")
	if err != nil{
			fmt.Println("数据库配置异常",err)
			return
	}
	Mysql_addr = appconf.String("mysql_addr")
	Mysql_prot = appconf.String("mysql_prot")
	Mysql_db = appconf.String("mysql_db")
	Mysql_name = appconf.String("mysql_name")
	Mysql_psw = appconf.String("mysql_psw")
	return
}

func init(){
	Inconfig()
}
