package set

import (
	"reflect"
	"testing"
)

func TestSet_IntersectEmpty(t *testing.T) {
	t.Run("intersect empty", func(t *testing.T) {
		nonEmpty := Of[int](1, 2)
		intersection := nonEmpty.Intersect(Empty[int]())
		if len(intersection) != 0 {
			t.Error("intersection with empty set must be the empty set")
		}
	})
}

func TestSet_Intersect(t *testing.T) {
	t.Run("intersect two non empty sets", func(t *testing.T) {
		set1 := Of[int](1, 2, 3, 4, 5, 6)
		set2 := Of[int](2, 6, 8)
		intersection := set1.Intersect(set2)
		want := Of[int](2, 6)
		if !reflect.DeepEqual(intersection, want) {
			t.Errorf("intersection must be %v but was %v", want, intersection)
		}
	})
}

func TestSet_Add(t *testing.T) {
	t.Run("add element", func(t *testing.T) {
		testSet := Empty[int]()
		element := 1
		if testSet.Contains(element) {
			t.Error("set must not contain element prior to adding it")
		}
		testSet.Add(element)
		if !testSet.Contains(element) {
			t.Error("set must contain element after adding it")
		}
	})
}

func TestSet_Remove(t *testing.T) {
	t.Run("remove element", func(t *testing.T) {
		element := 1
		testSet := Of[int](element)
		if !testSet.Contains(element) {
			t.Error("set must contain element prior to removing it")
		}
		testSet.Remove(element)
		if testSet.Contains(element) {
			t.Error("set must not contain element after removing it")
		}
	})
}

func TestSet_Pop(t *testing.T) {
	t.Run("pop element", func(t *testing.T) {
		element := 1
		testSet := Of[int](element)
		if !testSet.Contains(element) {
			t.Error("set must contain element prior to popping it")
		}
		poppedElement := testSet.Pop()
		if testSet.Contains(element) {
			t.Error("set must not contain element after popping it")
		}
		if !reflect.DeepEqual(poppedElement, element) {
			t.Error("popped element must be the element")
		}
	})
}

func TestSet_PopEmptySet(t *testing.T) {
	t.Run("pop from empty set", func(t *testing.T) {
		emptySet := Empty[int]()

		elem := emptySet.Pop()
		var zeroValue int
		if elem != zeroValue {
			t.Errorf("popping the empty set should return the zero value of the type")
		}
	})
}

func TestSet_Clone(t *testing.T) {
	t.Run("clone set", func(t *testing.T) {
		testSet := Of[int](1, 2, 3)
		clonedSet := testSet.Clone()

		if &clonedSet == &testSet {
			t.Error("cloned must be on a different memory address")
		}

		if !reflect.DeepEqual(testSet, clonedSet) {
			t.Error("cloned set must have the same elements as the src set")
		}
	})
}

func TestSet_Copy(t *testing.T) {
	t.Run("copy set", func(t *testing.T) {
		testSet := Of[int](1, 2, 3)
		copiedSet := Empty[int]()

		testSet.Copy(copiedSet)

		if !reflect.DeepEqual(testSet, copiedSet) {
			t.Error("copied set must have the same elements as the src set")
		}
	})
}

func TestSet_Difference(t *testing.T) {
	t.Run("difference set", func(t *testing.T) {
		numbers := Of[int](1, 2, 3, 4, 5)
		evenNumbers := Of[int](2, 4, 6)

		oddNumbers := numbers.Difference(evenNumbers)
		want := Of[int](1, 3, 5)

		if !reflect.DeepEqual(oddNumbers, want) {
			t.Errorf("%v \\ %v should be %v but was %v", numbers, evenNumbers, want, oddNumbers)
		}
	})
}

func TestSet_String(t *testing.T) {
	tests := []struct {
		testSet        Set[int]
		possibleValues []string
	}{
		{Of(1, 2), []string{"{ 1 2 }", "{ 2 1 }"}},
		{Of(1), []string{"{ 1 }"}},
		{Empty[int](), []string{"{ }"}},
	}

	for _, tt := range tests {
		t.Run("set to string", func(t *testing.T) {
			noMatch := true
			testSetString := tt.testSet.String()
			for _, want := range tt.possibleValues {
				if want == testSetString {
					noMatch = false
				}
			}
			if noMatch {
				t.Errorf("string must be any of %v but was %s", tt.possibleValues, testSetString)
			}
		})
	}
}

func TestSet_Slice(t *testing.T) {
	tests := []struct {
		testSet   Set[int]
		wantSlice []int
	}{
		{Of(1, 2, 3), []int{1, 2, 3}},
		{Of(1), []int{1}},
		{Empty[int](), []int{}},
	}

	for _, tt := range tests {
		t.Run("set to slice", func(t *testing.T) {
			testSetSlice := tt.testSet.Slice()
			if len(testSetSlice) != len(tt.testSet) {
				t.Errorf("slice must have the same length as the set")
			}
			for _, e := range testSetSlice {
				if !tt.testSet.Contains(e) {
					t.Errorf("slice elements must be in the set")
				}
			}
		})
	}
}
