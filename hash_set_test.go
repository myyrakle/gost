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
