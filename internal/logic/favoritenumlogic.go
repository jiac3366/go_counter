package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
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
	entry, err := l.svcCtx.CountMetaModel.FindOneByBusinessIdTypes(l.ctx, in.ArticleId, "FAVORITE")
	switch err {
	case nil:
		return &go_counter.FavoriteNumResponse{Count: entry.Count}, nil
	case sqlc.ErrNotFound:
		return &go_counter.FavoriteNumResponse{Count: 0}, nil
	default:
		return nil, err
	}

}
