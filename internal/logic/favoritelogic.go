package logic

//package main

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"go_counter/go_counter"
	"go_counter/internal/svc"
)

type FavoriteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteLogic {
	return &FavoriteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteLogic) Favorite(in *go_counter.FavoriteRequest) (*go_counter.FavoriteResponse, error) {
	l.svcCtx.FavoritesModel.FindOne()

	return &go_counter.FavoriteResponse{}, nil
}
