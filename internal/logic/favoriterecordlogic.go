package logic

import (
	"context"

	"go_counter/go_counter"
	"go_counter/internal/svc"
	"go_counter/model/mysql"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFavoriteRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteRecordLogic {
	return &FavoriteRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FavoriteRecordLogic) FavoriteRecord(in *go_counter.FavoriteRecordRequest) (*go_counter.FavoriteRecordResponse, error) {
	userId := in.UserId
	articleId := in.ArticleId
	if _, err := l.svcCtx.FavoritesRecordModel.Insert(
		l.ctx,
		&mysql.FavoritesRecord{
			UserId:    userId,
			ArticleId: articleId,
		},
	); err != nil {
		logx.Errorw("FavoritesModel.Insert failed", logx.Field("err", err.Error()))
		return &go_counter.FavoriteRecordResponse{Success: false}, nil
	}
	return &go_counter.FavoriteRecordResponse{Success: true}, nil
}
