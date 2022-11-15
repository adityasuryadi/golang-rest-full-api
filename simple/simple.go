package simple

import "errors"

type SimpleRepository struct {
	Error bool
}

type SimpleService struct {
	*SimpleRepository
}

// buat provider
/*
injector parameter,jika kita butuh parameter dinamis untuk dependency
tinggal tambah aja ke provider nya
*/
func NewSimpleRepository(isError bool) *SimpleRepository {
	return &SimpleRepository{
		// injector parameter
		Error: isError,
	}
}

// buat provider
/*
handle error google wire
*/
func NewSimpleService(repository *SimpleRepository) (*SimpleService, error) {
	if repository.Error {
		return nil, errors.New("failed create service")
	} else {
		return &SimpleService{
			SimpleRepository: repository,
		}, nil
	}

}
