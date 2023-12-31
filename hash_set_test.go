package gost

import "testing"

func Test_HashSet_IsSubset(t *testing.T) {
	t.Parallel()

	set1 := HashSetNew[I32]()
	set1.Insert(I32(1))
	set1.Insert(I32(2))

	set2 := HashSetNew[I32]()
	set2.Insert(I32(1))
	set2.Insert(I32(2))
	set2.Insert(I32(3))

	Assert(set1.IsSubset(set2))
	Assert(!set2.IsSubset(set1))
}

func Test_HashSet_IsSuperset(t *testing.T) {
	t.Parallel()

	set1 := HashSetNew[I32]()
	set1.Insert(I32(1))
	set1.Insert(I32(2))

	set2 := HashSetNew[I32]()
	set2.Insert(I32(1))
	set2.Insert(I32(2))
	set2.Insert(I32(3))

	Assert(set2.IsSuperset(set1))
	Assert(!set1.IsSuperset(set2))
}

func Test_HashSet_Intersection(t *testing.T) {
	t.Parallel()

	set1 := HashSetNew[I32]()
	set1.Insert(I32(1))
	set1.Insert(I32(2))
	set1.Insert(I32(5))

	set2 := HashSetNew[I32]()
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

func Test_HashSet_IsDisjoint(t *testing.T) {
	t.Parallel()

	set1 := HashSetNew[I32]()
	set1.Insert(I32(1))
	set1.Insert(I32(2))

	set2 := HashSetNew[I32]()
	set2.Insert(I32(3))
	set2.Insert(I32(4))

	Assert(set1.IsDisjoint(set2))
	Assert(set2.IsDisjoint(set1))

	set2.Insert(I32(2))
	Assert(!set1.IsDisjoint(set2))
	Assert(!set2.IsDisjoint(set1))
}

func Test_HashSet_Union(t *testing.T) {
	t.Parallel()

	set1 := HashSetNew[I32]()
	set1.Insert(I32(1))
	set1.Insert(I32(2))
	set1.Insert(I32(3))

	set2 := HashSetNew[I32]()
	set2.Insert(I32(3))
	set2.Insert(I32(4))

	union := set1.Union(set2)

	Assert(union.Len() == 4)
	Assert(union.Contains(I32(1)))
	Assert(union.Contains(I32(2)))
	Assert(union.Contains(I32(3)))
	Assert(union.Contains(I32(4)))
}

func Test_HashSet_SymmetricDifference(t *testing.T) {
	t.Parallel()

	set1 := HashSetNew[I32]()
	set1.Insert(I32(1))
	set1.Insert(I32(2))
	set1.Insert(I32(5))

	set2 := HashSetNew[I32]()
	set2.Insert(I32(1))
	set2.Insert(I32(2))
	set2.Insert(I32(3))

	diff := set1.SymmetricDifference(set2)

	Assert(diff.Len() == 2)
	Assert(diff.Contains(I32(3)))
	Assert(diff.Contains(I32(5)))
}
