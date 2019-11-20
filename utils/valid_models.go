package utils

import (
	"fmt"
	"reflect"
)

type Valid_models struct {

}

func (v*Valid_models) Valid_Models(data map[string]interface{},model interface{}) (err_msg string){
	st := reflect.TypeOf(model)
	for i := 0; i < st.NumField(); i++ {
		a:=st.Field(i)
		//fmt.Println(st.Field(i).Name)
		//fmt.Println(a)
		if a.Tag.Get("is_null") == "false"{
				fmt.Println(a.Tag.Get("json"))
				_,ok:=data[st.Field(i).Name]
				_,ok1:= data[a.Tag.Get("json")]
				if ok == false && ok1 == false{
					return "缺少"+st.Field(i).Name
				}
		}

	}
	return ""
}

