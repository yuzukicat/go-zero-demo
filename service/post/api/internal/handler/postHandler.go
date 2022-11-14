package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-demo/service/post/api/internal/logic"
	"go-demo/service/post/api/internal/svc"
	"go-demo/service/post/api/internal/types"
)

func postHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetPostByIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewPostLogic(r.Context(), svcCtx)
		resp, err := l.Post(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
