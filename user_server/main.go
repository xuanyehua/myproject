package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	_ "myproject/models"
	"myproject/user_server/handler"
	example "myproject/user_server/proto/example"
	img "myproject/user_server/proto/imgpoto"
	"myproject/user_server/subscriber"
	_ "myproject/utils"
)

func main() {
	// New Service
	//换成grpc传输
	service := grpc.NewService(
		micro.Name("go.micro.srv.user_server"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()
	logs.Error("server test Error")
	// Register Handler
	err := example.RegisterExampleHandler(service.Server(), new(handler.Example))
	fmt.Println(err)
	err = img.RegisterImgHandler(service.Server(), new(handler.Img))
	fmt.Println(err)
	// Register Struct as Subscriber
	err = micro.RegisterSubscriber("go.micro.srv.user_server", service.Server(), new(subscriber.Example))
	fmt.Println(err)
	// Register Function as Subscriber
	err = micro.RegisterSubscriber("go.micro.srv.user_server", service.Server(), subscriber.Handler)
	fmt.Println(err)
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
