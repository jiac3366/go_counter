package logic

//package main

import (
	"context"
	"go_counter/model/mysql"

	"go_counter/go_counter"
	"go_counter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
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

// 收藏请求
func (l *FavoriteLogic) Favorite(in *go_counter.FavoriteRequest) (*go_counter.FavoriteResponse, error) {
	userId := in.UserId
	articleId := in.ArticleId
	if _, err := l.svcCtx.FavoritesModel.Insert(
		l.ctx,
		&mysql.Favorites{
			UserId:    userId,
			ArticleId: articleId,
		},
	); err != nil {
		logx.Errorw("FavoritesModel.Insert failed", logx.Field("err", err.Error()))
		return &go_counter.FavoriteResponse{Success: false}, nil
	}

	return &go_counter.FavoriteResponse{Success: true}, nil
}
