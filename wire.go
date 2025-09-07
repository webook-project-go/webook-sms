//go:build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/webook-project-go/webook-sms/grpc"
	"github.com/webook-project-go/webook-sms/ioc"
	"github.com/webook-project-go/webook-sms/repository"
	"github.com/webook-project-go/webook-sms/repository/cache"
	"github.com/webook-project-go/webook-sms/service"
	"github.com/webook-project-go/webook-sms/service/provider"
)

var thirdPartyProvider = wire.NewSet(
	ioc.InitRedis,
	ioc.InitGrpcServer,
	ioc.InitEtcd,
)
var serviceSet = wire.NewSet(
	provider.NewSMSMemory,
	service.NewCodeService,
	repository.NewCodeRepository,
	cache.NewCodeCache,
)

func InitApp() *App {
	wire.Build(
		wire.Struct(new(App), "*"),
		thirdPartyProvider,
		serviceSet,
		grpc.New,
	)
	return nil
}
