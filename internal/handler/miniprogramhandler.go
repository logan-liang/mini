package handler

import (
	"net/http"

	"miniProgram/internal/logic"
	"miniProgram/internal/svc"
	"miniProgram/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func MiniProgramHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMiniProgramLogic(r.Context(), svcCtx)
		resp, err := l.MiniProgram(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
