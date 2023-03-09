package handler

import (
	"net/http"

	"github.com/lcclicheng/XFMcode/xfm_code/service/interface/api/internal/logic"
	"github.com/lcclicheng/XFMcode/xfm_code/service/interface/api/internal/svc"
	"github.com/lcclicheng/XFMcode/xfm_code/service/interface/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
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
