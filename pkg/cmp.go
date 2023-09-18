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

func (self Int) Cmp(rhs Int) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self Int8) Cmp(rhs Int8) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self Int16) Cmp(rhs Int16) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self Int32) Cmp(rhs Int32) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self Int64) Cmp(rhs Int64) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self Uint) Cmp(rhs Uint) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self Uint8) Cmp(rhs Uint8) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self Uint16) Cmp(rhs Uint16) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self Uint32) Cmp(rhs Uint32) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self Uint64) Cmp(rhs Uint64) Ordering {
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

func (self Int) Eq(rhs Int) Bool {
	return self == rhs
}

func (self Int8) Eq(rhs Int8) Bool {
	return self == rhs
}

func (self Int16) Eq(rhs Int16) Bool {
	return self == rhs
}

func (self Int32) Eq(rhs Int32) Bool {
	return self == rhs
}

func (self Int64) Eq(rhs Int64) Bool {
	return self == rhs
}

func (self Uint) Eq(rhs Uint) Bool {
	return self == rhs
}

func (self Uint8) Eq(rhs Uint8) Bool {
	return self == rhs
}

func (self Uint16) Eq(rhs Uint16) Bool {
	return self == rhs
}

func (self Uint32) Eq(rhs Uint32) Bool {
	return self == rhs
}

func (self Uint64) Eq(rhs Uint64) Bool {
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
