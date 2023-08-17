package logic

import (
	"context"
	"log"

	"go-zero-cloud-disk/core/internal/svc"
	"go-zero-cloud-disk/core/internal/types"
	"go-zero-cloud-disk/core/models"
	"go-zero-cloud-disk/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRepositorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRepositorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRepositorySaveLogic {
	return &UserRepositorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRepositorySaveLogic) UserRepositorySave(req *types.UserRepositorySaveRequest, userIdentity string) error {
	ur := &models.UserRepository{
		Identity:           utils.GenUUID(),
		UserIdentity:       userIdentity,
		ParentId:           req.ParentId,
		RepositoryIdentity: req.RepositoryIdentity,
		Ext:                req.Ext,
		Name:               req.Name,
	}
	result := l.svcCtx.MDB.Create(ur)
	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}
