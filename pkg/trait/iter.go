package trait

import "github.com/myyrakle/gost/pkg/option"

type Iterator[T any] interface {
	Next() option.Option[T]
}
