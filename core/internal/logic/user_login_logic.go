package logic

import (
	"context"
	"errors"

	"go-zero-cloud-disk/core/internal/svc"
	"go-zero-cloud-disk/core/internal/types"
	"go-zero-cloud-disk/core/models"
	"go-zero-cloud-disk/utils"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
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

	// find the first record with the given condition
	result := l.svcCtx.MDB.Where("name = ? AND password = ?", req.Name, req.Password).First(user)

	// return a RecordNotFound error if no record was found
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("incorrect name or password")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	token, err := utils.GenerateToken(user.Id, user.Identity, user.Name, 3600)
	if err != nil {
		return nil, err
	}

	resp = &types.LoginResponse{
		Token: token,
	}

	return resp, nil
}
