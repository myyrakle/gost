package gost

// A future represents an asynchronous computation obtained by use of async.
type Future[T any] interface {
	Await() T
}
