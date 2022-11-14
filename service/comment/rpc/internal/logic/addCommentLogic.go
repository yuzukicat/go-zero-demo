package logic

import (
	"context"

	"go-demo/service/comment/rpc/internal/svc"
	"go-demo/service/comment/rpc/types/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// -----------------------Comment-----------------------
func (l *AddCommentLogic) AddComment(in *comment.AddCommentReq) (*comment.AddCommentResp, error) {
	// todo: add your logic here and delete this line

	return &comment.AddCommentResp{}, nil
}
