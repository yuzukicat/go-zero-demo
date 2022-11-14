package logic

import (
	"context"

	"go-demo/service/post/rpc/internal/svc"
	"go-demo/service/post/rpc/types/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPostLogic {
	return &AddPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------Post-----------------------
func (l *AddPostLogic) AddPost(in *post.AddPostReq) (*post.AddPostResp, error) {
	// todo: add your logic here and delete this line

	return &post.AddPostResp{}, nil
}
