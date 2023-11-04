package gost

import "reflect"

func getTypeName[T any](value T) string {
	return reflect.TypeOf(value).Elem().Name()
}

func _Swap[T any](lhs *T, rhs *T) {
	*lhs, *rhs = *rhs, *lhs
}
