package logic

import (
	"context"

	"go-demo/service/post/rpc/internal/svc"
	"go-demo/service/post/rpc/types/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPostByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostByIdLogic {
	return &GetPostByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPostByIdLogic) GetPostById(in *post.GetPostByIdReq) (*post.GetPostByIdResp, error) {
	// todo: add your logic here and delete this line
	return &post.GetPostByIdResp{
		Post: &post.Post{
			Id:    "1",
			Title: "TestPost",
		},
	}, nil
}
