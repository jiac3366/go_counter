// Code generated by goctl. DO NOT EDIT.
// Source: go_counter.proto

package server

import (
	"context"

	"go_counter/go_counter"
	"go_counter/internal/logic"
	"go_counter/internal/svc"
)

type GoCounterServer struct {
	svcCtx *svc.ServiceContext
	go_counter.UnimplementedGoCounterServer
}

func NewGoCounterServer(svcCtx *svc.ServiceContext) *GoCounterServer {
	return &GoCounterServer{
		svcCtx: svcCtx,
	}
}

func (s *GoCounterServer) Ping(ctx context.Context, in *go_counter.PingRequest) (*go_counter.PingResponse, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

// 点赞请求
func (s *GoCounterServer) Like(ctx context.Context, in *go_counter.LikeRequest) (*go_counter.LikeResponse, error) {
	l := logic.NewLikeLogic(ctx, s.svcCtx)
	return l.Like(in)
}

// 收藏请求
func (s *GoCounterServer) Favorite(ctx context.Context, in *go_counter.FavoriteRequest) (*go_counter.FavoriteResponse, error) {
	l := logic.NewFavoriteLogic(ctx, s.svcCtx)
	return l.Favorite(in)
}

// 浏览请求
func (s *GoCounterServer) View(ctx context.Context, in *go_counter.ViewRequest) (*go_counter.ViewResponse, error) {
	l := logic.NewViewLogic(ctx, s.svcCtx)
	return l.View(in)
}
