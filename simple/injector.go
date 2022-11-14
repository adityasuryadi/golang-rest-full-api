//go:build wireinject
// +build wireinject

/*
penggunaan injector
injector sama seperti provider namun berisi config google wire
*/
package simple

import "github.com/google/wire"

func InitializedService() *SimpleService {
	/*
		kenapa return nil?karena semua nya kan di handle sama google wire
		dependency mana aja yang di butuhin
	*/
	wire.Build(NewSimpleRepository, NewSimpleService)

	return nil
}
