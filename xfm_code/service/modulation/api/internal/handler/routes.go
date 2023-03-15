// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/QueryCodeStatus",
				Handler: QueryCodeStatusHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/RequestConsumption",
				Handler: RequestConsumptionHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/OrderDetailsLogic",
				Handler: OrderDetailsHandler(serverCtx),
			},
		},
	)
}
