package logic

import (
	"api/internal/svc"
	"api/internal/types"
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("QueryConsumeStatusLogic_error", err)
		}
	}()

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
	conn, err := sql.Open("mysql", "root:123456@(127.0.0.1)/lccxfmcode?charset=utf8")
	fmt.Println("11111111111")
	if err != nil {
		fmt.Println("Failed to connect DB err:", err)
		return nil, errors.New("Failed to connect DB")
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		fmt.Println("ping-------------")
		return nil, err
	}

	//err = conn.QueryRow(StatusSql, params["uid"]).Scan(&resp.Status)
	rows, err := conn.Query(StatusSql, params["uid"])
	if err != nil {
		return nil, errors.New("Query DB Failed")
	}
	fmt.Println("2222222222")
	fmt.Println("sql:===", StatusSql)
	fmt.Println("uid:====", params)
	Status := ""
	fmt.Println(rows)
	for rows.Next() {
		fmt.Println("44444")
		err = rows.Scan(&Status)
		fmt.Println("555")
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("Status:%s", Status)
	}
	resp.Status = Status
	fmt.Println("3333333333333")
	resp.Code = 200
	resp.Errmsg = "success"
	fmt.Println("data:===", resp)
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
