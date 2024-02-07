// Code generated by goctl. DO NOT EDIT.
// Source: go_counter.proto

package gocounterclient

import (
	"context"

	"go_counter/go_counter"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = go_counter.Request
	Response = go_counter.Response

	GoCounter interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultGoCounter struct {
		cli zrpc.Client
	}
)

func NewGoCounter(cli zrpc.Client) GoCounter {
	return &defaultGoCounter{
		cli: cli,
	}
}

func (m *defaultGoCounter) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := go_counter.NewGoCounterClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
