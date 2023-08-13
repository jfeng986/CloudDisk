package logic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"go-zero-cloud-disk/core/internal/svc"
	"go-zero-cloud-disk/core/internal/types"
	"go-zero-cloud-disk/core/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type CoreLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCoreLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CoreLogic {
	return &CoreLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CoreLogic) Core(req *types.Request) (resp *types.Response, err error) {
	resp = new(types.Response)
	resp.Message = "hello"
	// get user list
	data := make([]*models.UserBasic, 0)
	err = models.Engine.Find(&data)
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
	}
	b, err := json.Marshal(data)
	if err != nil {
		log.Printf("json marshal error:%v", err)
	}
	dst := new(bytes.Buffer)
	err = json.Indent(dst, b, "", "\t")
	if err != nil {
		log.Printf("json indent error:%v", err)
	}
	fmt.Println(dst.String())
	resp = new(types.Response)
	resp.Message = dst.String()

	return
}
