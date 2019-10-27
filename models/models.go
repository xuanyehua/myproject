package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"myproject/utils"
)


type User struct {
	UserId   int    `orm:"column(user_id);pk"`
	Password string `orm:"column(password);size(128)"`
	UserName string `orm:"column(user_name);size(128)"`
	Phone    string `orm:"column(phone);size(11)"`
	Image    string `orm:"column(image);size(128)"`
}

func (t *User) TableName() string {
	return "user"
}


func init() {
	// 设置默认数据库

	err := orm.RegisterDataBase("default", "mysql", utils.Mysql_name+":"+utils.Mysql_psw+"@tcp("+utils.Mysql_addr+":"+utils.Mysql_prot+")/"+utils.Mysql_db+"?charset=utf8")
	if err != nil {
		fmt.Println("数据库异常",err)
	}
	orm.RegisterModel(new(User))
}
