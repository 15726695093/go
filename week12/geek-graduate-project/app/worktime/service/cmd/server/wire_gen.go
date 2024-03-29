// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"clock-in/app/worktime/service/internal/biz"
	"clock-in/app/worktime/service/internal/conf"
	"clock-in/app/worktime/service/internal/data"
	"clock-in/app/worktime/service/internal/server"
	"clock-in/app/worktime/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	workTimeRepo := data.NewWorkTimeRepo(dataData, logger)
	workTimeUsecase := biz.NewWorkTimeUsecase(workTimeRepo, logger)
	workTimeService := service.NewWorkTimeService(workTimeUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, workTimeService, logger)
	app := newApp(logger, grpcServer)
	return app, func() {
	}, nil
}
