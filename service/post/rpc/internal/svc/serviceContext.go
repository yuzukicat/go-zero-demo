package svc

import (
	"go-demo/service/post/model"
	"go-demo/service/post/rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	PostModel model.PostModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		PostModel: model.NewPostModel(conn),
	}
}
