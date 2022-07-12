// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tLinkFieldNames          = builder.RawFieldNames(&TLink{})
	tLinkRows                = strings.Join(tLinkFieldNames, ",")
	tLinkRowsExpectAutoSet   = strings.Join(stringx.Remove(tLinkFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	tLinkRowsWithPlaceHolder = strings.Join(stringx.Remove(tLinkFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheNavappTLinkIdPrefix   = "cache:navapp:tLink:id:"
	cacheNavappTLinkUuidPrefix = "cache:navapp:tLink:uuid:"
)

type (
	tLinkModel interface {
		Insert(ctx context.Context, data *TLink) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*TLink, error)
		FindOneByUuid(ctx context.Context, uuid string) (*TLink, error)
		Update(ctx context.Context, newData *TLink) error
		Delete(ctx context.Context, id int64) error
	}

	defaultTLinkModel struct {
		sqlc.CachedConn
		table string
	}

	TLink struct {
		Id        int64        `db:"id"`
		Uuid      string       `db:"uuid"`   // UUID
		Title     string       `db:"title"`  // 链接标题
		Desc      string       `db:"desc"`   // 链接描述
		Url       string       `db:"url"`    // 链接地址
		Status    int64        `db:"status"` // 资源状态:0默认|1被用户删除
		CreatedAt time.Time    `db:"created_at"`
		UpdatedAt sql.NullTime `db:"updated_at"`
		DeletedAt sql.NullTime `db:"deleted_at"`
	}
)

func newTLinkModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultTLinkModel {
	return &defaultTLinkModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`t_link`",
	}
}

func (m *defaultTLinkModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	navappTLinkIdKey := fmt.Sprintf("%s%v", cacheNavappTLinkIdPrefix, id)
	navappTLinkUuidKey := fmt.Sprintf("%s%v", cacheNavappTLinkUuidPrefix, data.Uuid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, navappTLinkIdKey, navappTLinkUuidKey)
	return err
}

func (m *defaultTLinkModel) FindOne(ctx context.Context, id int64) (*TLink, error) {
	navappTLinkIdKey := fmt.Sprintf("%s%v", cacheNavappTLinkIdPrefix, id)
	var resp TLink
	err := m.QueryRowCtx(ctx, &resp, navappTLinkIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tLinkRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTLinkModel) FindOneByUuid(ctx context.Context, uuid string) (*TLink, error) {
	navappTLinkUuidKey := fmt.Sprintf("%s%v", cacheNavappTLinkUuidPrefix, uuid)
	var resp TLink
	err := m.QueryRowIndexCtx(ctx, &resp, navappTLinkUuidKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `uuid` = ? limit 1", tLinkRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, uuid); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultTLinkModel) Insert(ctx context.Context, data *TLink) (sql.Result, error) {
	navappTLinkIdKey := fmt.Sprintf("%s%v", cacheNavappTLinkIdPrefix, data.Id)
	navappTLinkUuidKey := fmt.Sprintf("%s%v", cacheNavappTLinkUuidPrefix, data.Uuid)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?)", m.table, tLinkRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Uuid, data.Title, data.Desc, data.Url, data.Status, data.CreatedAt, data.UpdatedAt, data.DeletedAt)
	}, navappTLinkIdKey, navappTLinkUuidKey)
	return ret, err
}

func (m *defaultTLinkModel) Update(ctx context.Context, newData *TLink) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	navappTLinkIdKey := fmt.Sprintf("%s%v", cacheNavappTLinkIdPrefix, data.Id)
	navappTLinkUuidKey := fmt.Sprintf("%s%v", cacheNavappTLinkUuidPrefix, data.Uuid)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, tLinkRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Uuid, newData.Title, newData.Desc, newData.Url, newData.Status, newData.CreatedAt, newData.UpdatedAt, newData.DeletedAt, newData.Id)
	}, navappTLinkIdKey, navappTLinkUuidKey)
	return err
}

func (m *defaultTLinkModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheNavappTLinkIdPrefix, primary)
}

func (m *defaultTLinkModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", tLinkRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTLinkModel) tableName() string {
	return m.table
}
