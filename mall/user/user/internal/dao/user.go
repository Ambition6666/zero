package dao

import (
	"context"
	"fmt"
	"user/user/database"
	"user/user/internal/model"
)

type UserDao struct {
	conn *database.DBConn
}

func NewUserDao(conn *database.DBConn) *UserDao {
	return &UserDao{
		conn: conn,
	}
}

func (d *UserDao) Save(ctx context.Context, user *model.User) error {
	sql := fmt.Sprintf("insert into %s (name, gender) values (?, ?)", user.TableName())
	res, err := d.conn.Conn.ExecCtx(ctx, sql, user.Name, user.Gender)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = id
	return nil
}
