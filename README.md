# gost

![](https://img.shields.io/badge/language-Go-00ADD8) ![](https://img.shields.io/badge/version-v0.1.0-brightgreen) [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

Experience the true taste of Rust in Go

[document](https://pkg.go.dev/github.com/myyrakle/gost)

## Install

```
go get github.com/myyrakle/gost@{version}
```

## Example

```
package main

import (
	"math"

	"github.com/myyrakle/gost/pkg/option"
	"github.com/myyrakle/gost/pkg/primitive"
)

func CheckedAdd(a, b primitive.Int) option.Option[primitive.Int] {
	max := primitive.Int(math.MaxInt64)
	if (b > 0 && a > max-b) || (b < 0 && a < max-b) {
		return option.None[primitive.Int]()
	}

	return option.Some(a + b)
}

func main() {
	a := primitive.Int(1)
	b := primitive.Int(2)
	result := CheckedAdd(a, b)

	if result.IsSome() {
		println(result.Unwrap())
	} else {
		println("overflow")
	}

	a = primitive.Int(math.MaxInt64)
	b = primitive.Int(1)
	result = CheckedAdd(a, b)

	if result.IsSome() {
		println(result.Unwrap())
	} else {
		println("overflow")
	}
}
```
