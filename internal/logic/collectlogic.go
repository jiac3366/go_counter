package logic

import (
	"context"

	"go_counter/go_counter"
	"go_counter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CollectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCollectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CollectLogic {
	return &CollectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 收藏请求
func (l *CollectLogic) Collect(in *go_counter.FavoriteRequest) (*go_counter.FavoriteResponse, error) {
	// todo: add your logic here and delete this line

	return &go_counter.FavoriteResponse{}, nil
}
