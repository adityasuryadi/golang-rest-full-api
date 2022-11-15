// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package simple

import (
	"github.com/google/wire"
)

// Injectors from injector.go:

func InitializedService(isError bool) (*SimpleService, error) {
	simpleRepository := NewSimpleRepository(isError)
	simpleService, err := NewSimpleService(simpleRepository)
	if err != nil {
		return nil, err
	}
	return simpleService, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	databasePostgreSQL := NewDatabasePostgreSQL()
	databaseMongoDB := NewDatabaseMongoDB()
	databaseRepository := NewDataBaseRepository(databasePostgreSQL, databaseMongoDB)
	return databaseRepository
}

func InitializedFooBarService() *FooBarService {
	fooService := NewFooRepository()
	barService := NewBarService()
	fooBarService := NewFooBarService(fooService, barService)
	return fooBarService
}

// contoh struct provider
func InititalizedFooBar() *FooBar {
	foo := NewFoo()
	bar := NewBar()
	fooBar := &FooBar{
		Foo: foo,
		Bar: bar,
	}
	return fooBar
}

// injector.go:

// contoh provider set
var fooSet = wire.NewSet(NewFooRepository, NewFooService)

var barSet = wire.NewSet(NewBarRepository, NewBarService)
