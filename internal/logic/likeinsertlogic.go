package logic

import (
	"context"

	"go_counter/go_counter"
	"go_counter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeInsertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeInsertLogic {
	return &LikeInsertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LikeInsertLogic) LikeInsert(in *go_counter.LikeInsertRequest) (*go_counter.LikeInsertResponse, error) {
	// todo: add your logic here and delete this line

	return &go_counter.LikeInsertResponse{}, nil
}
