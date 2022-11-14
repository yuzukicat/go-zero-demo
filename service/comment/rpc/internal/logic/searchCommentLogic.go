package logic

import (
	"context"

	"go-demo/service/comment/rpc/internal/svc"
	"go-demo/service/comment/rpc/types/comment"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchCommentLogic {
	return &SearchCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchCommentLogic) SearchComment(in *comment.SearchCommentReq) (*comment.SearchCommentResp, error) {
	// todo: add your logic here and delete this line

	return &comment.SearchCommentResp{}, nil
}
