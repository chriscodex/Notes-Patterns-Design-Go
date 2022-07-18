package main

type topic interface {
	register(observer)
	broadcast()
}

type observer interface {
	getId() string
	updateValue(string)
}

type item struct {
	observers []observer
	name      string
	available bool
}

func newItem(name string) *item {
	return &item{
		name: name,
	}
}
