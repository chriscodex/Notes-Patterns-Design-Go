package main

type iProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type computer struct {
}
