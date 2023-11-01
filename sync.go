package gost

import "sync"

// A mutual exclusion lock.
//
//  locker := MutexNew[ISize](0)
//  ptr := locker.Lock()
//  *ptr = 1
//  locker.Unlock()
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
//
//  locker := MutexNew[ISize](0)
//  ptr := locker.Lock()
//  *ptr = 1
func (self *Mutex[T]) Lock() *T {
	self.lock.Lock()
	return self.value
}

// Immediately drops the guard, and consequently unlocks the mutex.
//
//  locker := MutexNew[ISize](0)
//  ptr := locker.Lock()
//  *ptr = 1
//  locker.Unlock()
func (self *Mutex[T]) Unlock() {
	self.lock.Unlock()
}

// An error returned by TryLock.
type TryLockError struct{}

func (self TryLockError) Error() string {
	return "Mutex is locked"
}

// Attempts to acquire this lock.
// If the lock could not be acquired at this time, then Err is returned. Otherwise, an RAII guard is returned. The lock will be unlocked when the guard is dropped.
//
//  locker := MutexNew[ISize](0)
//  ptr := locker.TryLock()
//  if ptr.IsErr() {
//      panic("Mutex is locked")
//  }
//  *ptr.Unwrap() = 1
func (self *Mutex[T]) TryLock() Result[*T] {
	if self.lock.TryLock() {
		return Ok[*T](self.value)
	} else {
		return Err[*T](TryLockError{})
	}
}
