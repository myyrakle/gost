package gost

type Ordering int

// Ordering enum values
const OrderingLess Ordering = Ordering(-1)
const OrderingEqual Ordering = Ordering(0)
const OrderingGreater Ordering = Ordering(1)

type Ord[T any] interface {
	Cmp(rhs T) Ordering
}

func (self Int) Cmp(rhs Int) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

type Eq[T any] interface {
	Eq(rhs T) Bool
}
