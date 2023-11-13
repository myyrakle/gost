package gost

import "testing"

func Test_BTreeSet_IsSubset(t *testing.T) {
	set1 := BTreeSetNew[I32]()
	set1.Insert(I32(1))
	set1.Insert(I32(2))

	set2 := BTreeSetNew[I32]()
	set2.Insert(I32(1))
	set2.Insert(I32(2))
	set2.Insert(I32(3))

	Assert(set1.IsSubset(set2))
	Assert(!set2.IsSubset(set1))
}

func Test_BTreeSet_IsSuperset(t *testing.T) {
	set1 := BTreeSetNew[I32]()
	set1.Insert(I32(1))
	set1.Insert(I32(2))

	set2 := BTreeSetNew[I32]()
	set2.Insert(I32(1))
	set2.Insert(I32(2))
	set2.Insert(I32(3))

	Assert(set2.IsSuperset(set1))
	Assert(!set1.IsSuperset(set2))
}
