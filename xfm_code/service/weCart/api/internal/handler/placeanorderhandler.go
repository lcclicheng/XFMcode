package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wecart/api/internal/logic"
	"wecart/api/internal/svc"
	"wecart/api/internal/types"
)

func PlaceAnOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PlaceAnOrderReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewPlaceAnOrderLogic(r.Context(), svcCtx)
		resp, err := l.PlaceAnOrder(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
