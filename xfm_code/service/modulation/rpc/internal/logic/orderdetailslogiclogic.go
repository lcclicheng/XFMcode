package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"modulation/rpc/internal/svc"
	"modulation/rpc/model"
	"modulation/rpc/modulation"
	"runtime/debug"
	"strings"
	"time"
	"tools/derror"
	"tools/dhttp"
	dmrdb "tools/mysql"
)

const DBURL = "root:12345@tcp(localhost:3306)/gozero?charset=utf8"

type OrderDetailsLogicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrderDetailsLogicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDetailsLogicLogic {
	return &OrderDetailsLogicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// OrderDetailsLogic 二维码订单查询
func (l *OrderDetailsLogicLogic) OrderDetailsLogic(in *modulation.OrderDetailsReq) (result *modulation.OrderDetailsResp, err error) {
	defer func() {
		if r := recover(); r != nil {
			_, err = derror.ResponsePanicErr(nil, r.(string))
			logx.Error(err, "OrderDetailsLogic_recover()", "in", in)
			debug.PrintStack()
		}
	}()
	prams := make(map[string]string, 2)
	prams["OutRequestNo"] = in.OutRequestNo
	prams["PayOutTradeNo"] = in.PayOutTradeNo
	where := ""
	args := []string{}

	if prams["OutRequestNo"] != "" {
		where += " AND OutRequestNo=? "
		args = append(args, prams["PayOutTradeNo"])

		delete(prams, prams["PayOutTradeNo"])

	} else if prams["PayOutTradeNo"] != "" {
		where += " AND PayOutTradeNo=? "
		args = append(args, prams["PayOutTradeNo"])

		delete(prams, prams["OutRequestNo"])

	}

	err = dhttp.ParamsCheck(prams)
	if err != nil {
		obj, err := derror.ResponseErr(nil, err)
		result = InterfaceToStruct(obj)
		return result, err
	}

	var userBuilderQueryRows = strings.Join(builder.RawFieldNames(&model.OrderDetails{}), ",")
	sql := "SELECT " + userBuilderQueryRows + " FROM orderDetails WHERE #{where} ORDER BY PayTime DESC LIMIT 1,10"

	cli := dmrdb.GetMysqlPoolCli(DBURL, 20)
	relData, err := cli.QuerySqlForMap(sql, 5*time.Second, args)
	if err != nil {
		logx.Error(err, "mysql查询错误", "args", args, "sql", sql)
		obj, err := derror.ResponseDBErr(nil, "DB failed to query")
		result = InterfaceToStruct(obj)
		return result, err
	}

	for _, v1 := range relData {
		for _, v2 := range result.Data {
			v2.PayStatus = v1["PayStatus"]
			v2.PayDate = v1["PayDate"]
			v2.PayTime = v1["PayTime"]
			v2.TotalFee = v1["TotalFee"]
			v2.PayCouponFee = v1["PayCouponFee"]
			v2.PayOutTradeNo = v1["PayOutTradeNo"]
			v2.PayErrDesc = v1["PayErrDesc"]
			v2.Uid = v1["Uid"]
			v2.PayType = v1["PayType"]
			v2.PayTypeTradeNo = v1["PayTypeTradeNo"]
			v2.OutRequestNo = v1["OutRequestNo"]
			v2.DimensionalCode = v1["DimensionalCode"]
			v2.BarCode = v1["BarCode"]
		}
	}
	obj, err := derror.ResponseSuccess(result)
	result = InterfaceToStruct(obj)
	return result, err
}

// InterfaceToStruct interface转任意类型，每个logic.go中返回类型`*modulation.OrderDetailsResp`记得替换成logic函数的返回类型
func InterfaceToStruct(object interface{}) *modulation.OrderDetailsResp {
	obj := object.(*modulation.OrderDetailsResp)
	return obj
}
