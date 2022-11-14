package svc

import (
	"go-demo/service/post/api/internal/config"
	"go-demo/service/post/rpc/post"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	PostRpc post.Post
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		PostRpc: post.NewPost(zrpc.MustNewClient(c.PostRpc)),
	}
}
