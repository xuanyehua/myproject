package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/micro/go-grpc"
	"net/http"
	"time"

	example "github.com/micro/examples/template/srv/proto/example"
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
