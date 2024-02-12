package logic

import (
	"context"

	"go_counter/go_counter"
	"go_counter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewInsertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewViewInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewInsertLogic {
	return &ViewInsertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ViewInsertLogic) ViewInsert(in *go_counter.ViewInsertRequest) (*go_counter.ViewInsertResponse, error) {
	// todo: add your logic here and delete this line

	return &go_counter.ViewInsertResponse{}, nil
}
