package main

import (
	"fmt"
	"github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"myproject/user_server/handler"
	"myproject/user_server/subscriber"

	_ "myproject/models"
	example "myproject/user_server/proto/example"
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

	// Register Handler
	err := example.RegisterExampleHandler(service.Server(), new(handler.Example))
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
