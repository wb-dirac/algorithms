package heap

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
)

type Int int

func (Int) MinKey() Comparable {
	return Int(math.MinInt64)
}

func (i Int) Compare(j Comparable) int {
	if i > j.(Int) {
		return 1
	} else if i == j {
		return 0
	} else {
		return -1
	}
}

func TestBuildHeap(t *testing.T) {
	var list = make([]Comparable, 100)
	for i := 0; i < 100; i++ {
		list[i] = Int(rand.Intn(10000))
	}
	heapIns := BuildHeap(list)
	heapIns.Println()
}

func TestHeapSort(t *testing.T) {
	var list = make([]Comparable, 100)
	for i := 0; i < 100; i++ {
		list[i] = Int(rand.Intn(10000))
	}
	HeapSort(list)
	for i := 0; i < 99; i++ {
		fmt.Print(list[i], ",")
		if list[i].Compare(list[i+1]) == 1 {
			t.Error("sort error")
		}
	}
}

var testList = []Comparable{Int(9), Int(100), Int(62), Int(55), Int(39), Int(90)}

func TestHeap_Max(t *testing.T) {
	heap := BuildHeap(testList)
	if heap.Max().Compare(Int(100)) != 0 {
		t.Errorf("heap.max fail")
	}
	if max, err := heap.ExtractMax(); err != nil {
		t.Error(err)
	} else if max.Compare(Int(100)) != 0 {
		t.Errorf("heap.ExtractMax return not max")
	} else {
		heap.Println()
		if heap.Size() != 5 {
			t.Errorf("heap.Size not right: %d", heap.Size())
		}
	}
}

func TestHeap_IncreaseKey(t *testing.T) {
	heap := BuildHeap(testList)
	heap.Println()
	_ = heap.IncreaseKey(4, Int(60))
	heap.Println()
	if heap.Get(1).Compare(Int(60)) != 0 {
		t.Errorf("Increase Failed, heap.Get(1): %d != 60", heap.Get(1))
	}
}

func TestHeap_Insert(t *testing.T) {
	heap := BuildHeap(testList)
	heap.Insert(Int(99))
	heap.Insert(Int(69))
	if heap.Get(1).Compare(Int(69)) != 0 {
		t.Errorf("test Insert fail")
	}
	if heap.Get(2).Compare(Int(99)) != 0 {
		t.Errorf("test Insert fail")
	}
	heap.Println()
}
