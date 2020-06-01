package sort

import "testing"

func TestQuickSort(t *testing.T) {
	testSort(t, QuickSort)
}

func TestMergeSort(t *testing.T) {
	testSort(t, MergeSort)
}

func TestInsertionSort(t *testing.T) {
	A := []int{1, 90, 37, 38, 9}
	InsertionSort(A)
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			t.Errorf("insertion sort error %v\n", A)
			return
		}
	}
	t.Logf("sort result: %v\n", A)
}

func testSort(t *testing.T, sortFunc func([]int, int, int)) {
	A := []int{1, 90, 37, 38, 9}
	sortFunc(A, 0, len(A)-1)
	for i := 0; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			t.Errorf("sort error %v\n", A)
			return
		}
	}
	t.Logf("sort result: %v\n", A)
}

