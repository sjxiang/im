package svc

import (
	"im/app/user/api/internal/config"
	"im/app/user/rpc/usersrv"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRPC usersrv.UserSrv  // gRPC client 更合适
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRPC: usersrv.NewUserSrv(zrpc.MustNewClient(c.UserRpc)),
	}
}
