package logic

import (
	"context"

	"go-demo/service/post/api/internal/svc"
	"go-demo/service/post/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostLogic {
	return &PostLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostLogic) Post(req *types.GetPostByIdReq) (resp *types.GetPostByIdResp, err error) {
	// todo: add your logic here and delete this line

	return
}
