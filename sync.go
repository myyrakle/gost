package gost

import "sync"

type Mutex[T any] struct {
	value *T
	lock  sync.Mutex
}

func MutexNew[T any](value T) Mutex[T] {
	return Mutex[T]{
		value: &value,
		lock:  sync.Mutex{},
	}
}

func (self *Mutex[T]) Lock() *T {
	self.lock.Lock()
	return self.value
}

func (self *Mutex[T]) Unlock() {
	self.lock.Unlock()
}
