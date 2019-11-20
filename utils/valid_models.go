package utils

import (
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
		if a.Tag.Get("is_null") == "fale"{
				_,ok:=data[st.Field(i).Name]
				if ok == false {
					return "缺少"+st.Field(i).Name
				}

		}

	}
	return ""
}

