package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im/app/user/api/internal/logic"
	"im/app/user/api/internal/svc"
)

// 获取用户信息
func detailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
