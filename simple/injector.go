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

func InitializedService() (*SimpleService, error) {
	/*
		kenapa return nil?karena semua nya kan di handle sama google wire
		dependency mana aja yang di butuhin
	*/
	wire.Build(NewSimpleRepository, NewSimpleService)

	// karena di atas tambah return error tambah return 1 nil lagi
	return nil, nil
}
