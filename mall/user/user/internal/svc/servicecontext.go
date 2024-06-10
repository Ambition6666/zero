package svc

import (
	"user/user/database"
	"user/user/internal/config"
	"user/user/internal/dao"
	"user/user/internal/repo"
)

type ServiceContext struct {
	Config   config.Config
	UserRepo repo.UserRepo
}

func NewServiceContext(c config.Config) *ServiceContext {
	connect := database.Connect(c.Mysql.DataSource, c.CacheRedis)
	return &ServiceContext{
		Config:   c,
		UserRepo: dao.NewUserDao(connect),
	}
}
