// Code generated by goctl. DO NOT EDIT.

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	likesFieldNames          = builder.RawFieldNames(&Likes{})
	likesRows                = strings.Join(likesFieldNames, ",")
	likesRowsExpectAutoSet   = strings.Join(stringx.Remove(likesFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	likesRowsWithPlaceHolder = strings.Join(stringx.Remove(likesFieldNames, "`article_id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	likesModel interface {
		Insert(ctx context.Context, data *Likes) (sql.Result, error)
		FindOne(ctx context.Context, articleId int64) (*Likes, error)
		Update(ctx context.Context, data *Likes) error
		Delete(ctx context.Context, articleId int64) error
	}

	defaultLikesModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Likes struct {
		LikeId    sql.NullInt64 `db:"like_id"`
		UserId    int64         `db:"user_id"`
		ArticleId int64         `db:"article_id"`
		CreatedAt time.Time     `db:"created_at"`
	}
)

func newLikesModel(conn sqlx.SqlConn) *defaultLikesModel {
	return &defaultLikesModel{
		conn:  conn,
		table: "`likes`",
	}
}

func (m *defaultLikesModel) Delete(ctx context.Context, articleId int64) error {
	query := fmt.Sprintf("delete from %s where `article_id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, articleId)
	return err
}

func (m *defaultLikesModel) FindOne(ctx context.Context, articleId int64) (*Likes, error) {
	query := fmt.Sprintf("select %s from %s where `article_id` = ? limit 1", likesRows, m.table)
	var resp Likes
	err := m.conn.QueryRowCtx(ctx, &resp, query, articleId)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultLikesModel) Insert(ctx context.Context, data *Likes) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, likesRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.LikeId, data.UserId, data.ArticleId)
	return ret, err
}

func (m *defaultLikesModel) Update(ctx context.Context, data *Likes) error {
	query := fmt.Sprintf("update %s set %s where `article_id` = ?", m.table, likesRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.LikeId, data.UserId, data.ArticleId)
	return err
}

func (m *defaultLikesModel) tableName() string {
	return m.table
}
