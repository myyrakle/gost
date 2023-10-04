package gost

import "reflect"

type Ordering int

// Ordering enum values
const OrderingLess Ordering = Ordering(-1)
const OrderingEqual Ordering = Ordering(0)
const OrderingGreater Ordering = Ordering(1)

type Ord[T any] interface {
	Cmp(rhs T) Ordering
}

func (self ISize) Cmp(rhs ISize) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self ISize8) Cmp(rhs ISize8) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self ISize16) Cmp(rhs ISize16) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self ISize32) Cmp(rhs ISize32) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self ISize64) Cmp(rhs ISize64) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self USize) Cmp(rhs USize) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self USize8) Cmp(rhs USize8) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self USize16) Cmp(rhs USize16) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self USize32) Cmp(rhs USize32) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self USize64) Cmp(rhs USize64) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self Float32) Cmp(rhs Float32) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self Float64) Cmp(rhs Float64) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self String) Cmp(rhs String) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self Bool) Cmp(rhs Bool) Ordering {
	if self == rhs {
		return OrderingEqual
	} else if self {
		return OrderingGreater
	} else {
		return OrderingLess
	}
}

func (self Byte) Cmp(rhs Byte) Ordering {
	if self == rhs {
		return OrderingEqual
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingLess
	}
}

func (self Rune) Cmp(rhs Rune) Ordering {
	if self == rhs {
		return OrderingEqual
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingLess
	}
}

func castToOrd[T any](value T) Option[Ord[T]] {
	reflectedValue := reflect.ValueOf(value)

	if ord, ok := reflectedValue.Interface().(Ord[T]); ok {
		return Some[Ord[T]](ord)
	} else {
		return None[Ord[T]]()
	}
}

type Eq[T any] interface {
	Eq(rhs T) Bool
}

func (self ISize) Eq(rhs ISize) Bool {
	return self == rhs
}

func (self ISize8) Eq(rhs ISize8) Bool {
	return self == rhs
}

func (self ISize16) Eq(rhs ISize16) Bool {
	return self == rhs
}

func (self ISize32) Eq(rhs ISize32) Bool {
	return self == rhs
}

func (self ISize64) Eq(rhs ISize64) Bool {
	return self == rhs
}

func (self USize) Eq(rhs USize) Bool {
	return self == rhs
}

func (self USize8) Eq(rhs USize8) Bool {
	return self == rhs
}

func (self USize16) Eq(rhs USize16) Bool {
	return self == rhs
}

func (self USize32) Eq(rhs USize32) Bool {
	return self == rhs
}

func (self USize64) Eq(rhs USize64) Bool {
	return self == rhs
}

func (self Float32) Eq(rhs Float32) Bool {
	return self == rhs
}

func (self Float64) Eq(rhs Float64) Bool {
	return self == rhs
}

func (self String) Eq(rhs String) Bool {
	return self == rhs
}

func (self Bool) Eq(rhs Bool) Bool {
	return self == rhs
}

func (self Byte) Eq(rhs Byte) Bool {
	return self == rhs
}

func (self Rune) Eq(rhs Rune) Bool {
	return self == rhs
}

func castToEq[T any](value T) Option[Eq[T]] {
	reflectedValue := reflect.ValueOf(value)

	if casted, ok := reflectedValue.Interface().(Eq[T]); ok {
		return Some[Eq[T]](casted)
	} else {
		return None[Eq[T]]()
	}
}
