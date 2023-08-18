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

type UserFileNameUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileNameUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileNameUpdateLogic {
	return &UserFileNameUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileNameUpdateLogic) UserFileNameUpdate(req *types.UserFileNameUpdateRequest, userIdentity string) error {
	// Check if the name already exists
	var cnt int64
	subquery := l.svcCtx.MDB.Model(&models.UserRepository{}).Where("identity = ?", req.Identity).Select("parent_id")
	err := l.svcCtx.MDB.Model(&models.UserRepository{}).
		Where("name = ? AND parent_id = ?", req.Name, gorm.Expr("(?)", subquery)).
		Count(&cnt).Error
	if err != nil {
		return err
	}
	if cnt > 0 {
		return errors.New("name already exists")
	}

	// Update the file name
	data := &models.UserRepository{Name: req.Name}
	err = l.svcCtx.MDB.Model(&models.UserRepository{}).
		Where("identity = ? AND user_identity = ?", req.Identity, userIdentity).
		Updates(data).Error
	if err != nil {
		return err
	}

	return nil
}
