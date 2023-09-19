package gost

import "reflect"

func getTypeName[T any](value T) string {
	return reflect.TypeOf(value).Elem().Name()
}
