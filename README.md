# gost

![](https://img.shields.io/badge/language-Go-00ADD8) ![](https://img.shields.io/badge/version-v0.2.0-brightgreen) [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

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
	"fmt"
	"math"

	gost "github.com/myyrakle/gost/pkg"
)

func CheckedAdd(a, b gost.Int) gost.Option[gost.Int] {
	max := gost.Int(math.MaxInt64)
	if (b > 0 && a > max-b) || (b < 0 && a < max-b) {
		return gost.None[gost.Int]()
	}

	return gost.Some(a + b)
}

func main() {
	a := gost.Int(1)
	b := gost.Int(2)
	result := CheckedAdd(a, b)

	fmt.Println(result.Display())
	if result.IsSome() {
		println(result.Unwrap())
	} else {
		println("overflow")
	}

	a = gost.Int(math.MaxInt64)
	b = gost.Int(1)
	result = CheckedAdd(a, b)

	fmt.Println(result.Display())
	if result.IsSome() {
		println(result.Unwrap())
	} else {
		println("overflow")
	}

	vec := gost.Vec[gost.Int]{}
	vec.Push(gost.Int(1))
	vec.Push(gost.Int(2))
	vec.Push(gost.Int(3))

	newVec := vec.IntoIter().Map(func(x gost.Int) gost.Int { return x * 2 }).CollectToVec()
	fmt.Println(newVec.Display())

	newVec.Push(gost.Int(7))
	foo := newVec.IntoIter().Fold(gost.Int(0), func(a, b gost.Int) gost.Int { return a + b })
	fmt.Println(foo)
}
```
