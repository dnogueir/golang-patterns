package main

type subject interface {
	register(obs observer)
	deRegister(obs observer)
	notifyAll()
}
