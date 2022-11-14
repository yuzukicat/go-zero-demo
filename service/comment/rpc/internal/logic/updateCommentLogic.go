package logic

import (
	"context"

	"go-demo/service/comment/rpc/internal/svc"
	"go-demo/service/comment/rpc/types/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentLogic {
	return &UpdateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCommentLogic) UpdateComment(in *comment.UpdateCommentReq) (*comment.UpdateCommentResp, error) {
	// todo: add your logic here and delete this line

	return &comment.UpdateCommentResp{}, nil
}
