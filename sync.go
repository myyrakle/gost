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

// Acquires a mutex, blocking the current thread until it is able to do so.
func (self *Mutex[T]) Lock() *T {
	self.lock.Lock()
	return self.value
}

// Immediately drops the guard, and consequently unlocks the mutex.
func (self *Mutex[T]) Unlock() {
	self.lock.Unlock()
}

type TryLockError struct{}

func (self TryLockError) Error() string {
	return "Mutex is locked"
}

// Attempts to acquire this lock.
// If the lock could not be acquired at this time, then Err is returned. Otherwise, an RAII guard is returned. The lock will be unlocked when the guard is dropped.
func (self *Mutex[T]) TryLock() Result[*T] {
	if self.lock.TryLock() {
		return Ok[*T](self.value)
	} else {
		return Err[*T](TryLockError{})
	}
}
