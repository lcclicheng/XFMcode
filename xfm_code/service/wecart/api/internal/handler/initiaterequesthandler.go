package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wecart/api/internal/logic"
	"wecart/api/internal/svc"
	"wecart/api/internal/types"
)

func InitiateRequestHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InitiateRequestReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewInitiateRequestLogic(r.Context(), svcCtx)
		resp, err := l.InitiateRequest(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
