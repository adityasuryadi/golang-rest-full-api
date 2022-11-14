package test

import (
	"belajar_golang_rest_full_api/simple"
	"fmt"
	"testing"
)

func TestSimpleService(t *testing.T) {
	simpleService := simple.InitializedService()
	fmt.Println(simpleService)
}
