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
	conn := sqlx.NewMysql(c.MysqlConn.DSN)
	return &ServiceContext{
		Config:         c,
		FavoritesModel: mysql.NewFavoritesModel(conn),
		ArticleModel:   mysql.NewArticleModel(conn),
		LikesModel:     mysql.NewLikesModel(conn),
	}
}
