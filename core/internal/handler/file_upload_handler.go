package handler

import (
	"crypto/md5"
	"errors"
	"fmt"
	"net/http"
	"path"

	"go-zero-cloud-disk/core/internal/logic"
	"go-zero-cloud-disk/core/internal/svc"
	"go-zero-cloud-disk/core/internal/types"
	"go-zero-cloud-disk/core/models"
	"go-zero-cloud-disk/utils"

	"github.com/zeromicro/go-zero/rest/httpx"
	"gorm.io/gorm"
)

func FileUploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FileUploadRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.Error(w, err)
			return
		}

		b := make([]byte, fileHeader.Size)
		_, err = file.Read(b)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum(b))
		rp := new(models.RepositoryPool)

		// check if file exist using GORM
		result := svcCtx.MDB.Where("hash = ?", hash).First(rp)
		if result.Error != nil {
			if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
				httpx.Error(w, result.Error)
				return
			}
		} else {
			httpx.OkJson(w, &types.FileUploadResponse{Identity: rp.Identity, Ext: rp.Ext, Name: rp.Name})
			return
		}

		// upload to aws
		AWSPath, err := utils.AWSUpload(r)
		if err != nil {
			httpx.Error(w, err)
			return
		}
		req.Name = fileHeader.Filename
		req.Ext = path.Ext(fileHeader.Filename)
		req.Size = fileHeader.Size
		req.Hash = hash
		req.Path = AWSPath

		l := logic.NewFileUploadLogic(r.Context(), svcCtx)
		resp, err := l.FileUpload(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
