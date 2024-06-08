package svc

import "user/user/internal/config"

type ServiceContext struct {
	Config config.Config
	
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
