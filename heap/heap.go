package heap

import (
	"errors"
	"fmt"
)

type Comparable interface {
	Compare(v2 Comparable) int
	MinKey() Comparable
}

//最大二叉堆
type Heap struct {
	items    []Comparable
	heapSize int
}

func (h *Heap) Size() int {
	return h.heapSize
}

func (h *Heap) Get(i int) Comparable {
	return h.items[i]
}

func (h *Heap) Max() Comparable {
	if h.heapSize < 1 {
		return nil
	}
	return h.items[0]
}

var HeapUnderflowErr = errors.New("Heap underflow")

func (h *Heap) ExtractMax() (Comparable, error) {
	if h.heapSize < 1 {
		return nil, HeapUnderflowErr
	}
	max := h.items[0]
	h.items[0] = h.items[h.heapSize-1]
	h.heapSize--
	h.maxHeapify(0)
	return max, nil
}

var NewKeyErr = errors.New("new key is smaller than current key")

func (h *Heap) IncreaseKey(i int, newKey Comparable) error {
	if h.items[i].Compare(newKey) == 1 {
		return NewKeyErr
	}
	h.items[i] = newKey
	for ; i > 0 && h.items[i].Compare(h.items[parent(i)]) == 1; i = parent(i) {
		exchangeItem(h.items, i, parent(i))
	}
	return nil
}

func (h *Heap) Insert(node Comparable) {
	h.heapSize++
	if len(h.items) < h.heapSize {
		h.items = append(h.items, nil)
	}
	h.items[h.heapSize-1] = node.MinKey()
	_ = h.IncreaseKey(h.heapSize-1, node)
}

//維護二叉堆
func (h *Heap) maxHeapify(i int) {
	largest := i
	l := left(i)
	r := right(i)
	if l < h.heapSize && h.items[l].Compare(h.items[i]) == 1 {
		largest = l
	}
	if r < h.heapSize && h.items[r].Compare(h.items[largest]) == 1 {
		largest = r
	}
	if largest != i {
		exchangeItem(h.items, i, largest)
		h.maxHeapify(largest)
	}
}

func (h *Heap) Println() {
	for i := 0; i < h.heapSize; i++ {
		fmt.Print(i, ":", h.items[i], ",")
	}
	fmt.Println()
}

func exchangeItem(list []Comparable, i, j int) {
	temp := list[i]
	list[i] = list[j]
	list[j] = temp
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

func parent(i int) int {
	return (i - 1) / 2
}

//構建一個二叉堆
func BuildHeap(list []Comparable) *Heap {
	h := &Heap{
		items:    list,
		heapSize: len(list),
	}
	for i := h.heapSize / 2; i >= 0; i-- {
		h.maxHeapify(i)
	}
	return h
}

//二叉堆原地排序
func HeapSort(list []Comparable) {
	h := BuildHeap(list)
	for i := h.heapSize - 1; i >= 0; i-- {
		exchangeItem(h.items, 0, i)
		h.heapSize--
		h.maxHeapify(0)
	}
}
