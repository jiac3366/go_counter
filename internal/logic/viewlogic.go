package logic

import (
	"context"

	"go_counter/go_counter"
	"go_counter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewLogic {
	return &ViewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 浏览请求
func (l *ViewLogic) View(in *go_counter.ViewRequest) (*go_counter.ViewResponse, error) {
	// todo: add your logic here and delete this line

	return &go_counter.ViewResponse{}, nil
}
