package trait

import "github.com/myyrakle/gost/pkg/primitive"

// Ordering enum values
const OrderingLess = primitive.Int(-1)
const OrderingEqual = primitive.Int(0)
const OrderingGreater = primitive.Int(1)

type Ord[T any] interface {
	Cmp(rhs T) primitive.Int
}

type Eq[T any] interface {
	Eq(rhs T) primitive.Bool
}
