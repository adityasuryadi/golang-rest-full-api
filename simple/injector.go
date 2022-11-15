//go:build wireinject
// +build wireinject

/*
penggunaan injector
injector sama seperti provider namun berisi config google wire
*/
package simple

import "github.com/google/wire"

/*
untuk handle error tambahkan return error pada initializesService

*/

func InitializedService(isError bool) (*SimpleService, error) {
	/*
		kenapa return nil?karena semua nya kan di handle sama google wire
		dependency mana aja yang di butuhin
	*/
	wire.Build(NewSimpleRepository, NewSimpleService)

	// karena di atas tambah return error tambah return 1 nil lagi
	return nil, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(NewDatabaseMongoDB, NewDatabasePostgreSQL, NewDataBaseRepository)
	return nil
}

// contoh provider set
var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

// contoh struct provider
func InititalizedFooBar() *FooBar {
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "*"))
	return nil
}
