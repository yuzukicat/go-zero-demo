package logic

import (
	"context"

	"go-demo/service/post/rpc/internal/svc"
	"go-demo/service/post/rpc/types/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelPostLogic {
	return &DelPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelPostLogic) DelPost(in *post.DelPostReq) (*post.DelPostResp, error) {
	// todo: add your logic here and delete this line

	return &post.DelPostResp{}, nil
}
