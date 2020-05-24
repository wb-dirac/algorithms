package main

import "math"

/**
插入排序, 输入数组A, 将A原地从小到大排序
*/
func InsertionSort(A []int) {
	//从第二个元素开始遍历, 遍历到 j 时, j以前的元素都排好序了
	for j := 1; j < len(A); j++ {
		//保存当前元素值
		key := A[j]
		i := j - 1
		//下面这个循环将所有j前面大于A[j]的元素都向后挪一格屁股
		for ; i > 0 && key < A[i]; i-- {
			A[i+1] = A[i]
		}
		//现在key应该保存到它应该呆的位置,注意这里为什么是A[i+1],而不是A[i]设置为key.
		//因为在上面的循环中最后一次i被减了1但不满足条件,所以应该加回来
		A[i+1] = key
	}
}

/**
归并排序, 输入数组A, 排序的前后范围p,r, .即函数对A中p到r的元素原地排序.
*/
func MergeSort(A []int, p, r int) {
	//多于一个元素,应用分治
	if p < r {
		//找到中间元素,比如A[0:10], 那么中间q=4
		q := (r + p) / 2
		MergeSort(A, p, q)
		MergeSort(A, q+1, r)
		//合并
		merge(A, p, q, r)
	}
}

func merge(A []int, p, q, r int) {
	//欲合并的两个数组的长度
	n1 := q - p + 1
	n2 := r - q
	//这里复制两个数组到L,R.
	//注意为了简化代码,这里数组长度都加了1, 把最后一个元素设为哨兵,这样做以防止两个数组的指标超出范围.
	//如果不这样做需要作更多判断.
	L := make([]int, n1+1)
	R := make([]int, n2+1)
	copy(L, A[p:q+1])
	copy(R, A[q+1:r+1])
	L[n1] = math.MaxInt64
	R[n2] = math.MaxInt64
	//i, j 分别指示L, R当前的位置
	i, j := 0, 0
	//比较当前L, R的元素,较小的压入A
	for k := p; k <= r; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}
}
