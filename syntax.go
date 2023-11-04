package gost

type ifElse[T any] struct {
	condition  Bool
	expression func() T
}

type ifElseContext[T any] struct {
	returnValue *T
}

// If-else statement.
//
//	n := I32(5)
//	foo := If(n == 10, func() String {
//		return "This is 10"
//	}).ElseIf(n == 5, func() String {
//		return "This is 5"
//	}).Else(func() String {
//	 	return "This is not 5 or 10"
//	})
//	AssertEq(foo, String("This is 5"))
func If[T any](condition Bool, expression func() T) ifElseContext[T] {
	context := ifElseContext[T]{}

	if condition {
		value := expression()
		context.returnValue = &value
	}

	return context
}

func (self ifElseContext[T]) Else(expression func() T) T {
	if self.returnValue != nil {
		return *self.returnValue
	}

	return expression()
}

func (self ifElseContext[T]) ElseIf(condition Bool, expression func() T) ifElseContext[T] {
	if condition {
		value := expression()
		self.returnValue = &value
	}

	return self
}
