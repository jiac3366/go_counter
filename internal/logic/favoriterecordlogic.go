package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"go_counter/go_counter"
	"go_counter/internal/svc"
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

func (l *FavoriteRecordLogic) FavoriteRecordV1(in *go_counter.FavoriteRecordRequest) (
	*go_counter.FavoriteRecordResponse, error) {
	// 有并发问题：

	//if _, err := l.svcCtx.FavoritesRecordModel.Insert(
	//	l.ctx,
	//	&mysql.FavoritesRecord{
	//		UserId:    in.UserId,
	//		ArticleId: in.ArticleId,
	//	},
	//); err != nil {
	//	logx.Errorw("FavoritesModel.Insert failed", logx.Field("err", err.Error()))
	//	return &go_counter.FavoriteRecordResponse{Success: false}, nil
	//}

	//entry, _ := l.svcCtx.CountMetaModel.FindOneByBusinessIdTypes(l.ctx, in.ArticleId, "FAVORITE")
	//if err != nil {
	//	...
	//}

	//entry.Count += 1
	//err := l.svcCtx.CountMetaModel.Update(
	//	l.ctx,
	//	entry,
	//)

	// v1 for update
	err := l.svcCtx.CountMetaModel.Incr(l.ctx, in.ArticleId, "FAVORITE", in.UserId)
	if err != nil {
		return &go_counter.FavoriteRecordResponse{Success: false}, err
	}

	return &go_counter.FavoriteRecordResponse{Success: true}, nil
}
