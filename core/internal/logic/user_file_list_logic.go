package logic

import (
	"context"
	"time"

	"go-zero-cloud-disk/core/define"
	"go-zero-cloud-disk/core/internal/svc"
	"go-zero-cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserFileListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserFileListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserFileListLogic {
	return &UserFileListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserFileListLogic) UserFileList(req *types.UserFileListRequest, userIdentity string) (resp *types.UserFileListResponse, err error) {
	uf := make([]*types.UserFile, 0)
	resp = new(types.UserFileListResponse)

	// Pagination
	size := req.Size
	if size <= 0 {
		size = define.PageSize
	}
	page := req.Page
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * size

	// Query the user file list from the database
	db := l.svcCtx.MDB.Table("user_repository").
		Where("parent_id = ? AND user_identity = ?", req.Identity, userIdentity).
		Select("user_repository.id, user_repository.identity, user_repository.repository_identity, "+
			"user_repository.ext, user_repository.name, repository_pool.path, repository_pool.size").
		Joins("LEFT JOIN repository_pool ON user_repository.repository_identity = repository_pool.identity").
		Where("user_repository.deleted_at = ? OR user_repository.deleted_at IS NULL", time.Time{}.Format(define.Datetime)).
		Limit(size).
		Offset(offset)

	if err := db.Find(&uf).Error; err != nil {
		return nil, err
	}

	// Query the total number of user files
	var cnt int64
	if err := l.svcCtx.MDB.Table("user_repository").
		Where("parent_id = ? AND user_identity = ?", req.Identity, userIdentity).
		Count(&cnt).Error; err != nil {
		return nil, err
	}

	resp.List = uf
	resp.Count = cnt
	return resp, nil
}
