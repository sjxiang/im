package svc

import (
	"im/app/user/api/internal/config"
	"im/app/user/rpc/usercenter"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRPC usercenter.UserCenter  // gRPC client 更合适
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRPC: usercenter.NewUserCenter(zrpc.MustNewClient(c.UserRpc)),
	}
}
