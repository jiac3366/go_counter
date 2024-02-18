package logic

import (
	"context"

	"go_counter/go_counter"
	"go_counter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteNumLogic {
	return &FavoriteNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteNumLogic) FavoriteNum(in *go_counter.FavoriteNumRequest) (*go_counter.FavoriteNumResponse, error) {
	// todo: add your logic here and delete this line

	return &go_counter.FavoriteNumResponse{}, nil
}
