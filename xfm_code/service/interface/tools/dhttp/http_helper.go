package dhttp

import (
	"strings"

	"git.woa.com/pdata/common/httpserver"
	"github.com/gin-gonic/gin"
)

// DataMoreRsp rsp通用返回
type DataMoreRsp struct {
	httpserver.HttpErrorCommon
	Data interface{} `json:"data"`
}

// GetParamsMap 获取URL参数同时获取Path的appIndex
func GetParamsMap(ginCtx *gin.Context) map[string]string {
	vars := httpserver.GetParamsMap(ginCtx)
	urlPathSlice := strings.Split(ginCtx.Request.URL.Path, "/")
	if len(urlPathSlice) >= 4 {
		vars["index"] = urlPathSlice[3]
	}
	return vars
}
