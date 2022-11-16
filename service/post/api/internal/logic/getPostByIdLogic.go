package logic

import (
	"context"

	"go-demo/service/post/api/internal/svc"
	"go-demo/service/post/api/internal/types"
	"go-demo/service/post/rpc/types/post"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPostByIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPostByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPostByIdLogic {
	return &GetPostByIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetPostByIdLogic) GetPostById(req *types.GetPostByIdReq) (resp *types.GetPostByIdResp, err error) {
	// todo: add your logic here and delete this line
	// IdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("Id")))
	// logx.Infof("Id: %s", IdNumber)
	// Id, err := IdNumber.Int64()
	// if err != nil {
	// 	return nil, err
	// }

	// use post rpc
	onePost, err := l.svcCtx.PostRpc.GetPostById(l.ctx, &post.GetPostByIdReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.GetPostByIdResp{
		Post: &types.Post{
			Id:        onePost.Post.Id,
			CreatedAt: onePost.Post.CreatedAt,
			UpdatedAt: onePost.Post.UpdatedAt,
			Title:     onePost.Post.Title,
			Published: onePost.Post.Published,
			Desc:      onePost.Post.Desc,
		},
	}, nil
}
