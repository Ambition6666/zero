package database

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type DBConn struct {
	Conn      sqlx.SqlConn
	ConnCache sqlc.CachedConn
}

func Connect(datasource string, conf cache.CacheConf) *DBConn {
	sqlConn := sqlx.NewMysql(datasource)
	cachedConn := sqlc.NewConn(sqlConn, conf)
	return &DBConn{
		Conn:      sqlConn,
		ConnCache: cachedConn,
	}
}
