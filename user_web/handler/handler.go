package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/astaxie/beego/logs"
	"github.com/micro/go-grpc"
	image_api "image"
	"image/png"
	"io/ioutil"
	common "myproject/user_server/proto/common"
	example "myproject/user_server/proto/example"
	image "myproject/user_server/proto/imgpoto"
	"net/http"
	"strconv"
	"time"
)

func ExampleCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//初始化grpc
	server := grpc.NewService()
	server.Init()
	fmt.Println("1")
	// call the backend service
	//grpc获取句柄
	exampleClient := example.NewExampleService("go.micro.srv.user_server",server.Client())
	rsp, err := exampleClient.Call(context.TODO(), &example.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func ImageCall(w http.ResponseWriter, r *http.Request) {
	//初始化grpc
	server := grpc.NewService()
	server.Init()
	image_cli:=image.NewImgService("go.micro.srv.user_server",server.Client())

	if r.Method == "GET" {
		err:=r.ParseForm()
		fmt.Println("get_img")
		fmt.Println(r.Form["uuid"][0])
		//fmt.Println(reflect.TypeOf(r.Form["uuid"][0]))
		uuid,err :=strconv.ParseInt(r.Form["uuid"][0],10,64)
		rsp,err :=image_cli.Call(context.TODO(),&image.Request{Uuid:uuid})
		if err != nil{
			logs.Error(err)
			return
		}
		fmt.Println(rsp)
		var image_new image_api.RGBA
		image_new.Stride = int(rsp.Stride)
		image_new.Rect.Min.X = int(rsp.Min.X)
		image_new.Rect.Min.Y = int(rsp.Min.Y)
		image_new.Rect.Max.Y = int(rsp.Man.Y)
		image_new.Rect.Max.X = int(rsp.Man.X)
		image_new.Pix = []uint8(rsp.Pix)
		var image1 captcha.Image
		image1.RGBA = &image_new
		err = png.Encode(w,image1)
		fmt.Println("++++++++",err)
		logs.Info("ok")
		logs.Error("aaaaaa")
	}
}

func LoginApi(w http.ResponseWriter, r *http.Request) {
	server := grpc.NewService()
	server.Init()
	common_cli := common.NewCommonService("go.micro.srv.user_server",server.Client())
	if r.Method == "POST"{
		//var post_param map[string]interface{}
		body,_ := ioutil.ReadAll(r.Body)
		//fmt.Println(string(body))
		rep,err :=common_cli.Call(context.TODO(),&common.Request{Data:string(body)})
		if err != nil{
			logs.Error(err)
		}
		fmt.Println(rep)
		var data map[string]interface{}
		data = make(map[string]interface{})
		data["code"] = rep.Code
		data["data"] = rep.Data
		data["msg"] = rep.Msg
		post_data,err := json.Marshal(data)
		if err != nil{
			logs.Error(err)
			return
		}
		_,err =fmt.Fprint(w,string(post_data))
		if err != nil{
			logs.Error(err)
			return
		}
	}
	if r.Method == "PUT"{
		fmt.Println("put")

	}
}


