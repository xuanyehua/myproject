package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"myproject/models"
	common "myproject/user_server/proto/common"
	"myproject/user_server/utils"
	utils2 "myproject/utils"
	"time"
)

type Login struct {

}


func (l *Login) Call(ctx context.Context, req *common.Request, rsp *common.Response) error{
	fmt.Println(req)
	var data map[string]interface{}
	err:=json.Unmarshal([]byte(req.Data),&data)
	if err != nil{
		logs.Error(err)
		return err
	}
	v := utils2.Valid_models{}
	err_msg := v.Valid_Models(data,models.RabcUser{})
	if err_msg != ""{
		rsp.Data = ""
		rsp.Code = 400
		rsp.Msg = err_msg
		return nil
	}
	validation_l := []string{"user_name","login_name","password","mobile","email","uuid","verification"}
	for _,i := range validation_l{
		fmt.Println(i)
		if utils.Map_in(i,data)== false {
			rsp.Data = ""
			rsp.Code = 400
			rsp.Msg = "缺少" + i
			return nil
		}
	}
	fmt.Println(data)
	var p utils.Password
	md5_psw,salt:=p.Encode_Password(data["password"].(string))
	var user models.RabcUser
	user.Password = md5_psw
	user.Salt = salt
	user.LoginName = data["login_name"].(string)
	user.UserName = data["user_name"].(string)
	user.Mobile = data["mobile"].(string)
	user.Email = data["email"].(string)
	user.CreateTime=time.Now()
	user.UpdateTime=time.Now()
	user.LastLoginTime=time.Now()
	o := orm.NewOrm()
	_,err =o.Insert(&user)
	if err != nil{
		logs.Error(err)
		fmt.Println(err)
		rsp.Data = ""
		rsp.Code = 400
		rsp.Msg = "保存失败"
		return nil
	}

	rsp.Data = ""
	rsp.Code = 200
	rsp.Msg = "ok"
	return nil
}








