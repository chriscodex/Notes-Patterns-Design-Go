package main

import "fmt"

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

func (i *item) updateAvailable() {
	fmt.Printf("Item %s is available\n", i.name)
	i.available = true
	i.broadcast()
}
