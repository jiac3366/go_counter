package logic

import (
	"context"

	"go_counter/go_counter"
	"go_counter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeNumLogic {
	return &LikeNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 点赞请求
func (l *LikeNumLogic) LikeNum(in *go_counter.LikeNumRequest) (*go_counter.LikeNumResponse, error) {
	// todo: add your logic here and delete this line

	return &go_counter.LikeNumResponse{}, nil
}
