package RBTree

import (
	"sort"
	"testing"
)

func TestTree(t *testing.T) {
	testValues := []int{81, 87, 47, 59, 18, 25, 40, 56, 0, 94, 11, 62, 89, 28, 74, 45, 37, 6, 95, 66, 58, 88, 90,
		15, 41, 8, 31, 29, 85, 26, 13, 63, 33, 78, 24, 53, 57, 21, 99, 5, 38, 3, 55, 51, 10, 61, 2, 83, 46,
		76, 77, 96, 20, 23, 43, 91, 36, 7, 52, 98}
	tree := Tree{}
	if len(tree.GetOrdered()) != 0 {
		t.FailNow()
	}
	for i := 0; i < len(testValues); i++ {
		tree.Put(testValues[i])
	}
	sort.Ints(testValues)
	ordered := tree.GetOrdered()
	if len(ordered) != len(testValues) {
		t.FailNow()
	}
	for i := 0; i < len(testValues); i++ {
		if ordered[i] != testValues[i] {
			t.FailNow()
		}
	}

	if tree.Get(11).Value != 11 {
		t.FailNow()
	}
	if tree.Get(999) != nil {
		t.FailNow()
	}
}
