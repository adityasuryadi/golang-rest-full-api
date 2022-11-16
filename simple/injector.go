//go:build wireinject
// +build wireinject

/*
penggunaan injector
injector sama seperti provider namun berisi config google wire
*/
package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

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

/* injector yang salah ketika binding interface
salaha karena type data nya beda walaupun NewSayHelloImpl karena balikan nya struct NewSayHelloImpl
karena NewHelloService return nya SayHello
*/
// func InititalizedHelloService() *HelloService {
// 	wire.Build(NewHelloService, NewSayHelloImpl)
// 	return nil
// }

// buat dulu menggunakan provider set
// jadi kita beritahu kan menggunakan binding jika ada yang butuh SayHelloImpl return SayHello
var helloSet = wire.NewSet(
	NewSayHelloImpl, wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializedHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}

// binding values
// kadang ada kasus dimana kita ingin melakukan dependency injection terhadap value yang sudah ada,tanpa membuat provider
var fooValue = &Foo{}
var barValue = &Bar{}

func InitializedFooBarUsingValue() *FooBar {
	// ketika di manapun ada yang butuh barValue atau FooValue maka akan di inject ke provider nya
	wire.Build(wire.Value(barValue), wire.Value(fooValue), wire.Struct(new(FooBar), "*"))
	return nil
}

// injector interface value
func InitalizedReader() io.Reader {
	// ketika ada yang butuh io reader maka value nya di inject os.stdin
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

// struct field provider
// membuat provider dari field dari struct
func InitializedConfiguration() *Configuration {
	// membuat sebuah provider dari sebuah field dari data si Application dan ambil field nya
	// google wire akan mencari di mana provider yang return value nya Application ambil data field configuration
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}

// clean up function
// jika provider membuat object yang membutuhkan proses clean up setelah object di buat,maka provider bisa mengembalikan closure

func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}
