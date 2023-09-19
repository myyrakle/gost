# gost

![](https://img.shields.io/badge/language-Go-00ADD8) ![](https://img.shields.io/badge/version-v0.4.0-brightgreen) [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

![](./etc/gorris.jpg)

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

	gost "github.com/myyrakle/gost"
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

	if result.IsSome() {
		gost.Println("result: {}", result.Unwrap())
	} else {
		gost.Println("result: overflow")
	}

	a = gost.Int(math.MaxInt64)
	b = gost.Int(1)
	result = CheckedAdd(a, b)

	if result.IsSome() {
		gost.Println("result: {}", result.Unwrap())
	} else {
		gost.Println("result: overflow")
	}

	vector := gost.Vec[gost.Int]{}
	vector.Push(gost.Int(3))
	vector.Push(gost.Int(1))
	vector.Push(gost.Int(2))
	vector.Push(gost.Int(4))
	vector.Sort()
	gost.Println("sorted Vec: {}", vector)

	newVec := vector.IntoIter().Map(func(x gost.Int) gost.Int { return x * 2 }).CollectToVec()
	gost.Println("mapped Vec: {}", newVec)

	newVec.Push(gost.Int(7))
	foo := newVec.IntoIter().Fold(gost.Int(0), func(a, b gost.Int) gost.Int { return a + b })
	gost.Println("fold value: {}", foo)

	hashMap := gost.HashMapNew[gost.String, gost.Int]()
	hashMap.Insert(gost.String("foo"), gost.Int(1))
	hashMap.Insert(gost.String("bar"), gost.Int(2))
	hashMap.Insert(gost.String("baz"), gost.Int(3))

	gost.Println("hashMap: {}", hashMap)
}
```
