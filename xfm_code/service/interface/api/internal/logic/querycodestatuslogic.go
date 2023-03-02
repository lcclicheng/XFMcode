package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/test/model"
	"strings"

	"api/internal/svc"
	"api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

//--查询消费码状态--

type QueryCodeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryCodeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryCodeStatusLogic {
	return &QueryCodeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryCodeStatusLogic) QueryCodeStatus(req *types.CodeStatusRequest) (resp *types.CodeStatusResponse, err error) {
	if len(strings.TrimSpace(req.Uid)) == 0 {
		return nil, errors.New("参数错误")
	}
	userInfo, err := l.svcCtx.UserModel.FindOneByMobile(req.Uid)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errors.New("用户信息不存在")
	default:
		return nil, err
	}

	if userInfo.Password != req.Uid {
		return nil, errors.New("用户密码不正确")
	}

	return
}
