//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"kratos_learn/wire/pkg"
)

func initService(c pkg.Config) (*pkg.Service, error) {
	panic(wire.Build(pkg.ProviderScheduleSet, pkg.ProviderDbSet, pkg.ProviderServiceSet))
}
