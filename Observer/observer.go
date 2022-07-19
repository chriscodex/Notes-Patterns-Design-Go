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

func (i *item) register(observer observer) {
	i.observers = append(i.observers, observer)
}

func (i *item) broadcast() {
	for _, observer := range i.observers {
		observer.updateValue(i.name)
	}
}

// Observer
type emailClient struct {
	id string
}

func (eC *emailClient) getId() string {
	return eC.id
}

func (eC *emailClient) updateValue(value string) {
	fmt.Printf("Sending email -%s available for client %s\n", value, eC.id)
}

func main() {
	nvidiaItem := newItem("RTX 3080")
	firstObserver := &emailClient{
		id: "12ab",
	}
	secondObserver := &emailClient{
		id: "34dc",
	}
	nvidiaItem.register(firstObserver)
	nvidiaItem.register(secondObserver)
	nvidiaItem.updateAvailable()
}
