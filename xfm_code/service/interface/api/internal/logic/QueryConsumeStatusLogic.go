package logic

import (
	"api/internal/svc"
	"api/internal/types"
	"context"
	"database/sql"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"regexp"
)

type QueryConsumeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryConsumeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryConsumeStatusLogic {
	return &QueryConsumeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryConsumeStatusLogic) QueryConsumeStatusLogic(req *types.QueryConsumeStatusReq) (resp *types.QueryConsumeStatusResp, err error) {
	params := make(map[string]string, 0)
	params["uid"] = req.Uid

	// 正则校验参数
	regOk, err := ParamsRegexp("^[0-9_a-z_A-Z\\,\\.\\&\\=#;_-]+$", params)
	if err != nil {
		return nil, errors.New("request params regexp err")

	}
	if !regOk {
		return nil, errors.New("request param regexp format is wrong")
	}

	StatusSql := "SELECT Status FROM codestatus WHERE Uid=?"
	conn, err := sql.Open("mysql", "root:123456@(127.0.0.1)/world")
	defer conn.Close()
	if err != nil {
		return nil, errors.New("Failed to connect DB")
	}
	rows, err := conn.Query(StatusSql, params["uid"])
	if err != nil {
		return nil, errors.New("Query DB Failed")
	}

	for rows.Next() {
		rows.Scan(&resp.Status)
	}

	return resp, nil
}

func ParamsRegexp(pattern string, params map[string]string) (bool, error) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return false, err
	}
	for _, v := range params {
		if !r.MatchString(v) {
			return false, nil
		}
	}
	return true, nil
}
