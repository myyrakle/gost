package trait

import "github.com/myyrakle/gost/pkg/primitive"

// A trait for converting a value to a String.
type ToString[T any] interface {
	ToString() primitive.String
}
