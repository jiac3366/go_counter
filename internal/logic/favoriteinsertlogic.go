package logic

import (
	"context"
	"go_counter/model/mysql"

	"go_counter/go_counter"
	"go_counter/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteInsertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteInsertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteInsertLogic {
	return &FavoriteInsertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteInsertLogic) FavoriteInsert(in *go_counter.FavoriteInsertRequest) (*go_counter.FavoriteInsertResponse, error) {

	// 插入记录
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
		return &go_counter.FavoriteInsertResponse{Success: false}, nil
	}

	return &go_counter.FavoriteInsertResponse{Success: true}, nil
}
