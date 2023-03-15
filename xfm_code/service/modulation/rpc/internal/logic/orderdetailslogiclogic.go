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

	for k, v := range relData {
		result.Data[k].PayStatus = v["PayStatus"]
		result.Data[k].PayDate = v["PayDate"]
		result.Data[k].PayTime = v["PayTime"]
		result.Data[k].TotalFee = v["TotalFee"]
		result.Data[k].PayCouponFee = v["PayCouponFee"]
		result.Data[k].PayOutTradeNo = v["PayOutTradeNo"]
		result.Data[k].PayErrDesc = v["PayErrDesc"]
		result.Data[k].Uid = v["Uid"]
		result.Data[k].PayType = v["PayType"]
		result.Data[k].PayTypeTradeNo = v["PayTypeTradeNo"]
		result.Data[k].OutRequestNo = v["OutRequestNo"]
		result.Data[k].DimensionalCode = v["DimensionalCode"]
		result.Data[k].BarCode = v["BarCode"]
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
