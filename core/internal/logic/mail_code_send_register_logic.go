package logic

import (
	"context"
	"errors"

	"go-zero-cloud-disk/core/internal/svc"
	"go-zero-cloud-disk/core/internal/types"
	"go-zero-cloud-disk/core/models"
	"go-zero-cloud-disk/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailCodeSendRegisterLogic {
	return &MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req *types.MailCodeSendRequest) error {
	cnt, err := models.Engine.Where("email = ?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return err
	}
	if cnt > 0 {
		err = errors.New("email already exist")
		return err
	}
	err = utils.MailSendCode(req.Email)
	if err != nil {
		return err
	}

	return nil
}
