//霍夫曼编码实现
package huffman

import (
	"container/heap"
	"fmt"
)

type hNode struct {
	freq        int
	left, right HuffmanTree
}

func (self hNode) Freq() int {
	return self.freq
}

type hLeaf struct {
	value rune
	freq  int
}

func (self hLeaf) Freq() int {
	return self.freq
}

type HuffmanTree interface {
	Freq() int
}

type PriorityQueue []HuffmanTree

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Freq() < pq[j].Freq()
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(node interface{}) {
	*pq = append(*pq, node.(HuffmanTree))
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func Decode(tree HuffmanTree, codes []byte) []rune {
	var data []rune
	node := tree
	n := len(codes)
	paddingSize := codes[n-1]
	ending := 0
	for index, code := range codes[0 : n-1] {
		if index == n-2 {
			ending = int(paddingSize)
		}
		for i := 7; i >= ending; i-- {
			nd := node.(hNode)
			if code&(1<<i) > 0 {
				node = nd.right
			} else {
				node = nd.left
			}
			if nd, ok := node.(hLeaf); ok {
				data = append(data, nd.value)
				node = tree
			}
		}
	}
	return data
}

type bitCode struct {
	code uint64
	bits byte
}

func getEncodeMap(tree HuffmanTree, encodeMap map[rune]bitCode, code uint64, bits byte) {
	switch n := tree.(type) {
	case hLeaf:
		encodeMap[n.value] = bitCode{code, bits}
	case hNode:
		bits++
		getEncodeMap(n.left, encodeMap, code<<1, bits)
		getEncodeMap(n.right, encodeMap, (code<<1)+1, bits)
	}
}

func codesAppend(codes []byte, remainPadding byte, code bitCode) ([]byte, byte) {
	if remainPadding == 0 {
		var newByte byte
		codes = append(codes, newByte)
		remainPadding = 8
	}
	n := len(codes)
	if code.bits > remainPadding {
		codes[n-1] |= byte(code.code>>int(code.bits-remainPadding)) & (1<<remainPadding - 1)
		code.bits -= remainPadding
		return codesAppend(codes, 0, code)
	} else {
		remainPadding -= code.bits
		codes[n-1] |= byte(code.code&(1<<code.bits-1)) << int(remainPadding)
		return codes, remainPadding
	}
}

func Encode(hTree HuffmanTree, raw []rune) []byte {
	encodeMap := make(map[rune]bitCode)
	getEncodeMap(hTree, encodeMap, 0, 0)

	var codes []byte
	var remainPadding byte
	for _, v := range raw {
		codes, remainPadding = codesAppend(codes, remainPadding, encodeMap[v])
	}
	//最后一个字节表示前面的字节有几个0填充
	codes = append(codes, remainPadding)
	return codes
}

func Println(tree HuffmanTree, prefix []byte) {
	switch n := tree.(type) {
	case hLeaf:
		fmt.Printf("%c(%d):%s\n", n.value, n.Freq(), prefix)
	case hNode:
		Println(n.left, append(prefix, '0'))
		Println(n.right, append(prefix, '1'))
	}
}

func HEncode(raw []rune) ([]byte, HuffmanTree) {
	charMapFreq := make(map[rune]int)
	for _, c := range raw {
		charMapFreq[c] += 1
	}
	ht := BuildHuffmanTree(charMapFreq)
	return Encode(ht, raw), ht
}

func BuildHuffmanTree(charFreqMap map[rune]int) HuffmanTree {
	pq := make(PriorityQueue, len(charFreqMap))
	var i int
	for c, freq := range charFreqMap {
		pq[i] = hLeaf{c, freq}
		i++
	}
	heap.Init(&pq)
	for pq.Len() > 1 {
		x, y := heap.Pop(&pq).(HuffmanTree), heap.Pop(&pq).(HuffmanTree)
		heap.Push(&pq, hNode{freq: x.Freq() + y.Freq(), left: x, right: y})
	}
	return heap.Pop(&pq).(HuffmanTree)
}
