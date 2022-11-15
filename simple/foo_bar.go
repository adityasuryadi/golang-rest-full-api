package simple

type Foo struct {
}

type Bar struct {
}

type FooBar struct {
	*Foo
	*Bar
}

func NewBar() *Bar {
	return &Bar{}
}

func NewFoo() *Foo {
	return &Foo{}
}
