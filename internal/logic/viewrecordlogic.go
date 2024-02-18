package logic

import (
	"context"

	"go_counter/go_counter"
	"go_counter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ViewRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewViewRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ViewRecordLogic {
	return &ViewRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ViewRecordLogic) ViewRecord(in *go_counter.ViewRecordRequest) (*go_counter.ViewRecordResponse, error) {
	// todo: add your logic here and delete this line

	return &go_counter.ViewRecordResponse{}, nil
}
