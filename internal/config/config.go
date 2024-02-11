package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	MysqlConn MysqlConn
}

type MysqlConn struct {
	DSN string
}
