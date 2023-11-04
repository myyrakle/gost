package gost

// Indicates whether a value is available or if the current task has been scheduled to receive a wakeup instead.
type Poll[T any] struct {
	value *T
}

// check if the value is ready
func (self *Poll[T]) IsReady() Bool {
	return self.value != nil
}

// check if the value is pending
func (self *Poll[T]) IsPending() Bool {
	return self.value == nil
}

// unwrap the value
func (self *Poll[T]) Unwrap() T {
	return *self.value
}

// A future represents an asynchronous computation obtained by use of async.
type Future[T any] interface {
	Await() T
	Poll() Poll[T]
}
