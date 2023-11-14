package gokio

import "github.com/myyrakle/gost"

type gokioFuture[T any] struct {
	ch chan T
}

func (self *gokioFuture[T]) Await() T {
	return <-self.ch
}

func Spawn[T any](f func() T) gost.Future[T] {
	ch := make(chan T)

	future := gokioFuture[T]{ch: ch}

	go func() {
		defer close(ch)
		ch <- f()
	}()

	return &future
}
