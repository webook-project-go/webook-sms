package main

import (
	"context"
	v1 "github.com/webook-project-go/webook-apis/gen/go/apis/sms/v1"
	_ "github.com/webook-project-go/webook-sms/config"
	"github.com/webook-project-go/webook-sms/ioc"
)

func main() {
	app := InitApp()
	shutdown := ioc.InitOTEL()
	defer shutdown(context.Background())
	v1.RegisterSMSServiceServer(app.Server, app.Service)
	err := app.Server.Serve()
	if err != nil {
		panic(err)
	}
}
