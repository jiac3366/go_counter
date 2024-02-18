package logic

import (
	"context"

	"go_counter/go_counter"
	"go_counter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewViewNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewNumLogic {
	return &ViewNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 浏览请求
func (l *ViewNumLogic) ViewNum(in *go_counter.ViewNumRequest) (*go_counter.ViewNumResponse, error) {
	// todo: add your logic here and delete this line

	return &go_counter.ViewNumResponse{}, nil
}
