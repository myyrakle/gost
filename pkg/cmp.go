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

type Eq[T any] interface {
	Eq(rhs T) Bool
}
