package logic

import (
	"context"

	"miniProgram/internal/svc"
	"miniProgram/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MiniProgramLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMiniProgramLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MiniProgramLogic {
	return &MiniProgramLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MiniProgramLogic) MiniProgram(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
