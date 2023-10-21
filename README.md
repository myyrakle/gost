# gost

![](https://img.shields.io/badge/language-Go-00ADD8) ![](https://img.shields.io/badge/version-v0.6.0-brightgreen) [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

![](./etc/gorris.jpg)

Experience the true taste of Rust in Go

[document](https://pkg.go.dev/github.com/myyrakle/gost)

## Install

```
go get github.com/myyrakle/gost@v0.6.0
```

## Example

```
package main

import (
	"math"

	gost "github.com/myyrakle/gost"
)

func CheckedAdd(a, b gost.ISize) gost.Option[gost.ISize] {
	max := gost.ISize(math.MaxInt)
	if (b > 0 && a > max-b) || (b < 0 && a < max-b) {
		return gost.None[gost.ISize]()
	}

	return gost.Some(a + b)
}

func main() {
	a := gost.ISize(1)
	b := gost.ISize(2)
	result := CheckedAdd(a, b)

	if result.IsSome() {
		gost.Println("result: {}", result.Unwrap())
	} else {
		gost.Println("result: overflow")
	}

	a = gost.ISize(math.MaxInt)
	b = gost.ISize(1)
	result = CheckedAdd(a, b)

	if result.IsSome() {
		gost.Println("result: {}", result.Unwrap())
	} else {
		gost.Println("result: overflow")
	}

	vector := gost.Vec[gost.ISize]{}
	vector.Push(gost.ISize(3))
	vector.Push(gost.ISize(1))
	vector.Push(gost.ISize(2))
	vector.Push(gost.ISize(4))
	vector.Sort()
	gost.Println("sorted Vec: {}", vector)

	newVec := vector.IntoIter().Map(func(x gost.ISize) gost.ISize { return x * 2 }).CollectToVec()
	gost.Println("mapped Vec: {}", newVec)

	newVec.Push(gost.ISize(7))
	foo := newVec.IntoIter().Fold(gost.ISize(0), func(a, b gost.ISize) gost.ISize { return a + b })
	gost.Println("fold value: {}", foo)

	hashMap := gost.HashMapNew[gost.String, gost.ISize]()
	hashMap.Insert(gost.String("foo"), gost.ISize(1))
	hashMap.Insert(gost.String("bar"), gost.ISize(2))
	hashMap.Insert(gost.String("baz"), gost.ISize(3))

	gost.Println("hashMap: {}", hashMap)

	linkedList := gost.LinkedListNew[gost.ISize]()
	linkedList.PushBack(gost.ISize(1))
	linkedList.PushFront(gost.ISize(2))
	linkedList.PushBack(gost.ISize(3))
	linkedList2 := gost.LinkedListNew[gost.ISize]()
	linkedList2.PushBack(gost.ISize(4))
	linkedList2.PushFront(gost.ISize(5))
	linkedList.Append(&linkedList2)

	gost.Println("linkedList: {}", linkedList)
}
```
