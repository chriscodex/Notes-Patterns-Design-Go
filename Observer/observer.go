package main

type topic interface {
	register(observer)
	broadcast()
}
