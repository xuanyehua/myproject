package utils

import (
	"fmt"
	"github.com/astaxie/beego/logs"
)


func init(){
	err :=logs.SetLogger(logs.AdapterMultiFile, `{"filename": "logs/log", "separate":["error", "warning", "info"]}`)
	if err != nil{
		fmt.Println(err)
	}
	logs.EnableFuncCallDepth(true)
	logs.SetLogFuncCallDepth(3)
	logs.Info("program start...")
}
