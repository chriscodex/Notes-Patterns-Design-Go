package main

type iProduct interface {
	setStock(stock int)
	getStock() int
}
