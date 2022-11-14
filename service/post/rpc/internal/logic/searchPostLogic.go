package logic

import (
	"context"

	"go-demo/service/post/rpc/internal/svc"
	"go-demo/service/post/rpc/types/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchPostLogic {
	return &SearchPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchPostLogic) SearchPost(in *post.SearchPostReq) (*post.SearchPostResp, error) {
	// todo: add your logic here and delete this line

	return &post.SearchPostResp{}, nil
}
