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

func Test_BTreeSet_Intersection(t *testing.T) {
	set1 := BTreeSetNew[I32]()
	set1.Insert(I32(1))
	set1.Insert(I32(2))
	set1.Insert(I32(5))

	set2 := BTreeSetNew[I32]()
	set2.Insert(I32(1))
	set2.Insert(I32(2))
	set2.Insert(I32(3))

	intersection := set1.Intersection(set2)

	Assert(intersection.Len() == 2)
	Assert(intersection.Contains(I32(1)))
	Assert(intersection.Contains(I32(2)))
	Assert(!intersection.Contains(I32(3)))
	Assert(!intersection.Contains(I32(5)))
}

func Test_BTreeSet_IsDisjoint(t *testing.T) {
	set1 := BTreeSetNew[I32]()
	set1.Insert(I32(1))
	set1.Insert(I32(2))

	set2 := BTreeSetNew[I32]()
	set2.Insert(I32(3))
	set2.Insert(I32(4))

	Assert(set1.IsDisjoint(set2))
	Assert(set2.IsDisjoint(set1))

	set2.Insert(I32(2))
	Assert(!set1.IsDisjoint(set2))
	Assert(!set2.IsDisjoint(set1))
}
