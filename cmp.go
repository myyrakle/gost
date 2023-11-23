package gost

import "reflect"

type Ordering int

// Ordering enum values
const OrderingLess Ordering = Ordering(-1)
const OrderingEqual Ordering = Ordering(0)
const OrderingGreater Ordering = Ordering(1)

// Trait for types that form a total order.
type Ord[T any] interface {
	Cmp(rhs T) Ordering
	Eq[T]
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

func (self I8) Cmp(rhs I8) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self I16) Cmp(rhs I16) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self I32) Cmp(rhs I32) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self I64) Cmp(rhs I64) Ordering {
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

func (self U8) Cmp(rhs U8) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self U16) Cmp(rhs U16) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self U32) Cmp(rhs U32) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self U64) Cmp(rhs U64) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self F32) Cmp(rhs F32) Ordering {
	if self < rhs {
		return OrderingLess
	} else if self > rhs {
		return OrderingGreater
	} else {
		return OrderingEqual
	}
}

func (self F64) Cmp(rhs F64) Ordering {
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

func (self Char) Cmp(rhs Char) Ordering {
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

// Trait for equality comparisons which are equivalence relations.
type Eq[T any] interface {
	Eq(rhs T) Bool
}

func (self ISize) Eq(rhs ISize) Bool {
	return self == rhs
}

func (self I8) Eq(rhs I8) Bool {
	return self == rhs
}

func (self I16) Eq(rhs I16) Bool {
	return self == rhs
}

func (self I32) Eq(rhs I32) Bool {
	return self == rhs
}

func (self I64) Eq(rhs I64) Bool {
	return self == rhs
}

func (self USize) Eq(rhs USize) Bool {
	return self == rhs
}

func (self U8) Eq(rhs U8) Bool {
	return self == rhs
}

func (self U16) Eq(rhs U16) Bool {
	return self == rhs
}

func (self U32) Eq(rhs U32) Bool {
	return self == rhs
}

func (self U64) Eq(rhs U64) Bool {
	return self == rhs
}

func (self U128) Eq(rhs U128) Bool {
	return self.high == rhs.high && self.low == rhs.low
}

func (self F32) Eq(rhs F32) Bool {
	return self == rhs
}

func (self F64) Eq(rhs F64) Bool {
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

func (self Char) Eq(rhs Char) Bool {
	return self == rhs
}

func castToEq[T any](value T) _InternalOption[Eq[T]] {
	reflectedValue := reflect.ValueOf(value)

	if casted, ok := reflectedValue.Interface().(Eq[T]); ok {
		return _Some[Eq[T]](casted)
	} else {
		return _None[Eq[T]]()
	}
}
