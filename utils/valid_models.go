package utils

import (
	"github.com/astaxie/beego/orm"
	"reflect"
)

type Valid_models struct {

}

func (v*Valid_models) Valid_Models(data map[string]interface{},model interface{}) (err_msg string){
	st := reflect.TypeOf(model)
	for i := 0; i < st.NumField(); i++ {
		a:=st.Field(i)
		if a.Tag.Get("is_null") == "false"{
				_,ok:=data[st.Field(i).Name]
				_,ok1:= data[a.Tag.Get("json")]
				if ok == false && ok1 == false{
					return "缺少"+st.Field(i).Name
				}
		}
		if a.Tag.Get("only") == "true" {
			b:=v.Valid_Only(st.Field(i).Name,data[a.Tag.Get("json")],model)
			if b == true{
				return st.Field(i).Name+"已存在"
			}
		}
	}
	return ""
}

func (v*Valid_models) Valid_Only(field string,value interface{},model interface{}) bool{
	o := orm.NewOrm()
	a :=o.QueryTable("rabc_user").Filter(field,value).Exist()
	return a
}