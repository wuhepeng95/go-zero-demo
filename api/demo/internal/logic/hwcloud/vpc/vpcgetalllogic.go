package vpc

import (
	"context"

	"demo/internal/svc"
	"demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VpcGetAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVpcGetAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VpcGetAllLogic {
	return &VpcGetAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VpcGetAllLogic) VpcGetAll(req *types.VpcGetAllReq) (resp *types.VpcGetAllResp, err error) {
	// todo: add your logic here and delete this line

	return
}
