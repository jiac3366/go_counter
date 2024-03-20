package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_counter/internal/config"
	"go_counter/model/mysql"
)

type ServiceContext struct {
	Config               config.Config
	ArticleModel         mysql.ArticleModel
	FavoritesRecordModel mysql.FavoritesRecordModel
	LikesRecordModel     mysql.LikesRecordModel
	CountMetaModel       mysql.CountMetaModel
	SqlxConn             sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MysqlConf.DSN)

	return &ServiceContext{
		SqlxConn:             mysqlConn,
		Config:               c,
		ArticleModel:         mysql.NewArticleModel(mysqlConn, c.RedisConf),
		FavoritesRecordModel: mysql.NewFavoritesRecordModel(mysqlConn, c.RedisConf),
		LikesRecordModel:     mysql.NewLikesRecordModel(mysqlConn, c.RedisConf),
		CountMetaModel:       mysql.NewCountMetaModel(mysqlConn, c.RedisConf),
	}
}
