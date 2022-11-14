package test

import (
	"belajar_golang_rest_full_api/simple"
	"fmt"
	"testing"
)

// pemanggilan dependency injection yg sudah di generate google wire

//tambah handle error
func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializedService()
	fmt.Println(err)
	fmt.Println(simpleService)
}
