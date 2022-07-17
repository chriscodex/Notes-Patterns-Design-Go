package main

import "fmt"

type iProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

// Computer
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

func (c computer) getStock() int {
	return c.stock
}

func (c computer) getName() string {
	return c.name
}

// Laptop
type laptop struct {
	computer
}

func newLaptop() iProduct {
	return &laptop{
		computer: computer{
			name:  "Laptop Computer",
			stock: 25,
		},
	}
}

// Desktop
type desktop struct {
	computer
}

func newDesktop() iProduct {
	return &desktop{
		computer: computer{
			name:  "Desktop Computer",
			stock: 35,
		},
	}
}

// Factory
func getComputerFactory(computerType string) (iProduct, error) {
	switch computerType {
	case "laptop":
		return newLaptop(), nil
	case "desktop":
		return newDesktop(), nil
	default:
		return nil, fmt.Errorf("Invalid computer type")
	}
}

func printNameAndStock(p iProduct) {

}
