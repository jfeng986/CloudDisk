package handler

import (
	"net/http"

	"go-zero-cloud-disk/core/internal/logic"
	"go-zero-cloud-disk/core/internal/svc"
	"go-zero-cloud-disk/core/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserFileNameUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserFileNameUpdateRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserFileNameUpdateLogic(r.Context(), svcCtx)
		err := l.UserFileNameUpdate(&req, r.Header.Get("UserIdentity"))
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
