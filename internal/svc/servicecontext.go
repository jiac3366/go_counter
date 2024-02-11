package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go_counter/internal/config"
	"go_counter/model/mysql"
)

type ServiceContext struct {
	Config         config.Config
	FavoritesModel mysql.FavoritesModel
	ArticleModel   mysql.ArticleModel
	LikesModel     mysql.LikesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysqlConn := sqlx.NewMysql(c.MysqlConf.DSN)

	return &ServiceContext{
		Config:         c,
		FavoritesModel: mysql.NewFavoritesModel(mysqlConn, c.RedisConf),
		ArticleModel:   mysql.NewArticleModel(mysqlConn, c.RedisConf),
		LikesModel:     mysql.NewLikesModel(mysqlConn, c.RedisConf),
	}
}
