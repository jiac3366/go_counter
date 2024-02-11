package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	MysqlConf MysqlConf
	RedisConf cache.CacheConf
}

type MysqlConf struct {
	DSN string
}
