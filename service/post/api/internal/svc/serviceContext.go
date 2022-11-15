package svc

import (
	"go-demo/service/post/api/internal/config"
	"go-demo/service/post/rpc/service"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	PostRpc service.Service
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		PostRpc: service.NewService(zrpc.MustNewClient(c.PostRpc)),
	}
}
