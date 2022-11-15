package logic

import (
	"context"
	"errors"

	"go-demo/service/post/model"
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
	onePost, err := l.svcCtx.PostModel.FindOne(l.ctx, in.Id)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errors.New("post does not exist")
	default:
		return nil, err
	}
	return &post.GetPostByIdResp{
		Post: &post.Post{
			Id:        onePost.Id,
			CreatedAt: onePost.CreatedAt.Unix(),
			UpdatedAt: onePost.UpdatedAt.Unix(),
			Title:     onePost.Title,
			Published: onePost.Published,
			Desc:      onePost.Desc,
		},
	}, nil
}
