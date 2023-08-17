package logic

import (
	"context"
	"errors"
	"time"

	"go-zero-cloud-disk/core/internal/svc"
	"go-zero-cloud-disk/core/internal/types"
	"go-zero-cloud-disk/core/models"
	"go-zero-cloud-disk/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEmailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmailCodeSendRegisterLogic {
	return &EmailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EmailCodeSendRegisterLogic) EmailCodeSendRegister(req *types.MailCodeSendRequest) error {
	var cnt int64
	err := l.svcCtx.MDB.Model(&models.UserBasic{}).Where("email = ?", req.Email).Count(&cnt).Error
	if err != nil {
		return err
	}
	if cnt > 0 {
		err = errors.New("email already exist")
		return err
	}

	code, err := utils.EmailSendCode(req.Email)
	if err != nil {
		return err
	}
	// request timeout
	// err = models.RDB.Set(l.ctx, req.Email, code, 60*time.Second).Err()

	err = l.svcCtx.RDB.Set(context.Background(), req.Email, code, 300*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}
