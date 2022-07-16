package main

type iProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type computer struct {
	name  string
	stock int
}

func (c computer) setStock(stock int) {
	c.stock = stock
}

func (c computer) setName(name string) {
	c.name = name
}
