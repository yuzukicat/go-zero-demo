package logic

import (
	"context"

	"go-demo/service/comment/rpc/internal/svc"
	"go-demo/service/comment/rpc/types/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentByIdLogic {
	return &GetCommentByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentByIdLogic) GetCommentById(in *comment.GetCommentByIdReq) (*comment.GetCommentByIdResp, error) {
	// todo: add your logic here and delete this line

	return &comment.GetCommentByIdResp{}, nil
}
