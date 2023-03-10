package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"modulation/api/internal/logic"
	"modulation/api/internal/svc"
	"modulation/api/internal/types"
)

func OrderDetailsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderDetailsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewOrderDetailsLogic(r.Context(), svcCtx)
		resp, err := l.OrderDetails(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
