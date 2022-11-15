package simple

type FooRepository struct {
}

type FooService struct {
	*FooService
}

func NewFooRepository() *FooService {
	return &FooService{}
}

func NewFooService() *FooRepository {
	return &FooRepository{}
}
