package logic

import (
	"context"

	"go-zero-cloud-disk/core/internal/svc"
	"go-zero-cloud-disk/core/internal/types"
	"go-zero-cloud-disk/core/models"
	"go-zero-cloud-disk/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	user := new(models.UserBasic)

	has, err := l.svcCtx.Engine.Where("name = ? AND password = ?", req.Name, req.Password).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, err
	}

	token, err := utils.GenerateToken(user.Id, user.Identity, user.Name, 3600)
	resp = new(types.LoginResponse)
	resp.Token = token

	return resp, nil
}
