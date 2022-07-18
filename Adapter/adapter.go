package main

import "fmt"

type payment interface {
	pay()
}

type cashPayment struct{}

func (cashPayment) pay() {
	fmt.Println("Payment using Cash")
}

func proccessPayment(p payment) {
	p.pay()
}

type bankPayment struct{}

func (bankPayment) pay(bankAccount int) {
	fmt.Printf("Paying using bankaccount %d\n", bankAccount)
}

type bankPaymentAdapter struct {
	bankPayment *bankPayment
	bankAccount int
}

func (bpa *bankPaymentAdapter) pay() {
	bpa.bankPayment.pay(bpa.bankAccount)
}

func main() {
	cash := &cashPayment{}
	proccessPayment(cash)
	bpa := &bankPaymentAdapter{
		bankPayment: &bankPayment{},
		bankAccount: 9889889,
	}
	proccessPayment(bpa)
}
