//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
)

func InisializeService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

func InisializeDatabaseRepository() *DatabaseRepository {
	wire.Build(NewDatabasePosgreeSQL, NewDatabaseMongoDB, NewDatabaseRepository)
	return nil
}
