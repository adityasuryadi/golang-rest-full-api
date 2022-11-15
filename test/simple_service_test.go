package test

import (
	"belajar_golang_rest_full_api/simple"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// pemanggilan dependency injection yg sudah di generate google wire

//tambah handle error
func TestSimpleService(t *testing.T) {
	simpleService, err := simple.InitializedService(true)
	fmt.Println(err)
	fmt.Println(simpleService)
}

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializedService(true)
	assert.NotNil(t, err)
	assert.Nil(t, simpleService)
}

func TestSimpeServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializedService(false)
	assert.NotNil(t, simpleService)
	assert.Nil(t, err)
}
