package logic

import (
	"context"

	"go_counter/go_counter"
	"go_counter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeRecordLogic {
	return &LikeRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LikeRecordLogic) LikeRecord(in *go_counter.LikeRecordRequest) (*go_counter.LikeRecordResponse, error) {
	// todo: add your logic here and delete this line

	return &go_counter.LikeRecordResponse{}, nil
}
