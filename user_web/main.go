package main

import (
        "github.com/astaxie/beego/logs"
        "github.com/micro/go-log"
        "net/http"

        "github.com/micro/go-web"
        "myproject/user_web/handler"
        _ "myproject/utils"
)

func main() {
	// create new web service
        service := web.NewService(
                web.Name("go.micro.web.user_web"),
                web.Version("latest"),
                web.Address("127.0.0.1:8080"),
        )
        logs.Error("web test Error")
	// initialise service
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/example/call", handler.ExampleCall)

	service.HandleFunc("/get_img", handler.ImageCall)

	service.HandleFunc("/login", handler.LoginApi)




	// run service
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
