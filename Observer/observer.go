package main

type topic interface {
	register(observer)
	broadcast()
}

type observer interface {
	getId() string
	updateValue(string)
}
