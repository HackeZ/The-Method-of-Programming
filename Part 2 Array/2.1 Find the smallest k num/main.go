package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	// Method 1
	arrayNum := []int{2, 7, 5, 8, 4, 1, 3}
	fmt.Println("Method 1 --> ", FindtheSmallest1(arrayNum, 3))
	// Method 2
	arrayNum = []int{2, 7, 5, 8, 4, 1, 3}
	fmt.Println("Method 2 --> ", FindtheSmallest2(arrayNum, 3))
	// Method 3
	arrayNum = []int{2, 7, 5, 8, 4, 1, 3}
	fmt.Printf("Method 3 -->  [")
	FindtheSmallest3(arrayNum, 3)
}

// Author: HackerZ
// Time: 2016/06/11 14:43

// Look for the smallest number of K
// Input:  [1, 2, 3, 4, 5, 6], 3
// Output: 1, 2, 3

// FindtheSmallest1 ALL SORT
// 解法1：全部排序法
func FindtheSmallest1(arrn []int, k int) []int {
	for i := 0; i < len(arrn); i++ {
		for j := 0; j < len(arrn)-i-1; j++ {
			if arrn[j] > arrn[j+1] {
				arrn[j], arrn[j+1] = arrn[j+1], arrn[j]
			}
		}
	}
	return arrn[0:k]
}

// FindtheSmallest2 Partial SORT
// 解法2：部分排序法
func FindtheSmallest2(arrn []int, k int) []int {
	sortedArr := arrn[0:k]
	sort.Ints(sortedArr)
	for i := k; i < len(arrn); i++ {
		if sortedArr[k-1] > arrn[i] {
			sortedArr[k-1] = arrn[i]
			sort.Ints(sortedArr)
		}
	}
	return sortedArr
}

// FindtheSmallest3 Heap instead of array
// 解法3：用堆代替数组

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push IntHeap push into heap
func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

// Pop IntHeap Pop the bigest int in Heap
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// FindtheSmallest3 Heap instead of array
func FindtheSmallest3(arrn []int, k int) {
	h := IntHeap(arrn)
	heap.Init(&h)
	for i := 0; i < k; i++ {
		fmt.Printf("%d ", heap.Pop(&h))
	}
	fmt.Printf("\b]")
}
