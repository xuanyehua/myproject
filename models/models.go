package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"myproject/utils"
	"time"
)


type RabcGroup struct {
	Id         int       `orm:"column(g_id);pk" description:"组id"`
	GroupName  string    `orm:"column(group_name);size(64)" description:"组名"`
	ParentGId  int       `orm:"column(parent_g_id)" description:"父id"`
	CreateTime time.Time `orm:"column(create_time);type(datetime)" description:"创建时间"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime)" description:"修改时间"`
	Desc       string    `orm:"column(desc);size(128)" description:"描述"`
}

func (t *RabcGroup) TableName() string {
	return "rabc_group"
}

type RabcUser struct {
	Id            int       `orm:"column(u_id);pk" description:"用户id" json:"u_id" `
	GId           int       `orm:"column(g_id)" description:"分组id" json:"g_id"`
	LoginName     string    `orm:"column(login_name);size(64)" description:"登录名" json:"login_name" is_null:"false"`
	Password      string    `orm:"column(password);size(64)" description:"密码" json:"password" is_null:"false"`
	UserName      string    `orm:"column(user_name);size(64)" description:"用户名" json:"user_name" is_null:"false"`
	Mobile        string    `orm:"column(mobile);size(20)" description:"手机号" json:"mobile" is_null:"false"`
	Email         string    `orm:"column(email);size(64)" description:"邮箱" json:"email" is_null:"false"`
	Salt         string    `orm:"column(salt);size(11)" description:"盐" json:"salt" is_null:"false"`
	CreateTime    time.Time `orm:"column(create_time);type(datetime)" json:"create_time" is_null:"false"`
	UpdateTime    time.Time `orm:"column(update_time);type(datetime)" description:"修改时间" json:"update_time" is_null:"false"`
	LastLoginTime time.Time `orm:"column(last_login_time);type(datetime)" json:"last_login_time" is_null:"false"`
	LoginCount    int       `orm:"column(login_count)" description:"登录次数" json:"login_count" is_null:"false"`
}

func (t *RabcUser) TableName() string {
	return "rabc_user"
}

type RbacGroupRole struct {
	Id  int `orm:"column(gr_id);pk"`
	GId int `orm:"column(g_id)" description:"组id"`
	RId int `orm:"column(r_id);null" description:"角色id"`
}

func (t *RbacGroupRole) TableName() string {
	return "rbac_group_role"
}

type RbacLog struct {
	Id         int       `orm:"column(log_id);pk"`
	Content    string    `orm:"column(content);size(128)" description:"操作内容"`
	UId        int       `orm:"column(u_id)" description:"操作人"`
	CreateTime time.Time `orm:"column(create_time);type(datetime)" description:"操作时间"`
}

func (t *RbacLog) TableName() string {
	return "rbac_log"
}

type RbacRight struct {
	Id         int    `orm:"column(ri_id);pk" description:"权限id"`
	ParentRiId int    `orm:"column(parent_ri_id)" description:"父权限"`
	RightName  string `orm:"column(right_name);size(64)" description:"权限名称"`
	RightDesc  string `orm:"column(right_desc);size(128)" description:"权限描述"`
}

func (t *RbacRight) TableName() string {
	return "rbac_right"
}

type RbacRole struct {
	Id         int       `orm:"column(r_id);pk" description:"角色id"`
	ParentRId  int       `orm:"column(parent_r_id)" description:"父角色id"`
	RoleName   string    `orm:"column(role_name);size(64)" description:"角色名称"`
	CreateTime time.Time `orm:"column(create_time);type(datetime)" description:"创建时间"`
	UpdateTime time.Time `orm:"column(update_time);type(datetime)"`
	Desc       string    `orm:"column(desc);size(128)" description:"角色描述"`
}

func (t *RbacRole) TableName() string {
	return "rbac_role"
}


type RbacRoleRight struct {
	Id   int `orm:"column(rr_id);pk"`
	RId  int `orm:"column(r_id)" description:"角色id"`
	RiId int `orm:"column(ri_id)" description:"权限id"`
}

func (t *RbacRoleRight) TableName() string {
	return "rbac_role_right"
}


func init() {
	// 设置默认数据库

	err := orm.RegisterDataBase("default", "mysql", utils.Mysql_name+":"+utils.Mysql_psw+"@tcp("+utils.Mysql_addr+":"+utils.Mysql_prot+")/"+utils.Mysql_db+"?charset=utf8")
	if err != nil {
		fmt.Println("数据库异常",err)
	}
	orm.RegisterModel(new(RabcGroup),new(RabcUser),new(RbacGroupRole),new(RbacLog),new(RbacRight),new(RbacRole),new(RbacRoleRight))
}
