package logic

import (
	"context"
	"errors"

	"go-zero-cloud-disk/core/internal/svc"
	"go-zero-cloud-disk/core/internal/types"
	"go-zero-cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDetailLogic {
	return &UserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDetailLogic) UserDetail(req *types.UserDetailRequest) (resp *types.UserDetailResponse, err error) {
	resp = &types.UserDetailResponse{}
	ub := new(models.UserBasic)

	result := l.svcCtx.MDB.Where("identity = ?", req.Identity).First(&ub)

	// return a RecordNotFound error if no record was found
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("user not exist")
	}

	if result.Error != nil {
		return nil, result.Error
	}

	resp.Name = ub.Name
	resp.Email = ub.Email
	return
}
