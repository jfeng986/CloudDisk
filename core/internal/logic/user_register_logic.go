package logic

import (
	"context"
	"errors"
	"log"

	"go-zero-cloud-disk/core/internal/svc"
	"go-zero-cloud-disk/core/internal/types"
	"go-zero-cloud-disk/core/models"
	"go-zero-cloud-disk/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	code, err := l.svcCtx.RDB.Get(l.ctx, req.Email).Result()
	if err != nil {
		return nil, err
	}
	if code != req.Code {
		err = errors.New("code is not correct")
		return nil, err
	}

	// check if name already exist
	var cnt int64
	err = l.svcCtx.MDB.Model(&models.UserBasic{}).Where("name = ?", req.Name).Count(&cnt).Error
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		err = errors.New("name already exist")
		return nil, err
	}

	user := &models.UserBasic{
		Identity: utils.GenUUID(),
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	}
	result := l.svcCtx.MDB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	log.Println("insert user: ID", user.Id)

	token, err := utils.GenerateToken(user.Id, user.Identity, user.Name, 3600)
	if err != nil {
		return nil, err
	}

	resp = &types.UserRegisterResponse{
		Token: token,
	}

	return resp, nil
}
