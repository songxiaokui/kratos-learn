// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"kratos_learn/wire/pkg"
)

// Injectors from wire.go:

func initService(c pkg.Config) (*pkg.Service, error) {
	database := pkg.NewDB(c)
	schedule := pkg.NewSchedule(c)
	service := pkg.NewService(database, schedule)
	return service, nil
}
