package simple

type SimpleRepository struct {
}

type SimpleService struct {
	*SimpleRepository
}

// buat provider
func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

// buat provider
func NewSimpleService(repository *SimpleRepository) *SimpleService {
	return &SimpleService{
		SimpleRepository: repository,
	}
}
