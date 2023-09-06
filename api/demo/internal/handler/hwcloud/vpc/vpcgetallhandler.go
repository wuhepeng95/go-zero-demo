package vpc

import (
	"net/http"

	"demo/internal/logic/hwcloud/vpc"
	"demo/internal/svc"
	"demo/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func VpcGetAllHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.VpcGetAllReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := vpc.NewVpcGetAllLogic(r.Context(), svcCtx)
		resp, err := l.VpcGetAll(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
