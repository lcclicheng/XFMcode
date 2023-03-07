package handler

import (
	"github.com/lcclicheng/XFMcode/xfm_code/service/interface/api/internal/logic"
	"github.com/lcclicheng/XFMcode/xfm_code/service/interface/api/internal/svc"
	"github.com/lcclicheng/XFMcode/xfm_code/service/interface/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ExpandHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ExpandReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewExpandLogic(r.Context(), svcCtx)
		resp, err := l.Expand(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
