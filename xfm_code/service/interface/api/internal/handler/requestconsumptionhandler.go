package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"interface/api/internal/logic"
	"interface/api/internal/svc"
	"interface/api/internal/types"
)

func RequestConsumptionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RequestConsumptionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRequestConsumptionLogic(r.Context(), svcCtx)
		resp, err := l.RequestConsumption(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
