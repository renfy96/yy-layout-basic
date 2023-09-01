//go:build wireinject
// +build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	repo "github.com/renfy96/yy-layout-v1/internal/repository"
	"github.com/renfy96/yy-layout-v1/internal/router"
	"github.com/renfy96/yy-layout-v1/pkg/log"
	"github.com/spf13/viper"
)

// RepositorySet 仓储
var RepositorySet = wire.NewSet(
	repo.NewDb,
)

// 请求服务
var repoSvc = wire.NewSet()

// SvcSet 业务服务
var SvcSet = wire.NewSet()

var ServerSet = wire.NewSet(router.NewServerHTTP)

func newApp(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		RepositorySet,
		repoSvc,
		SvcSet,
		ServerSet,
	))
}
