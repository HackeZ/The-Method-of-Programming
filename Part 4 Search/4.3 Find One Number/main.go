package main

import (
	"fmt"
	"sort"
)

func main() {
	// Method 1
	fmt.Println("=== Method 1 ===")
	Input := []int{0, 1, 1, 2, 1, 3}
	fmt.Println("Output --> ", SortedAndFind(Input))
	fmt.Println("=== Method 1 ===")

	fmt.Println("")

	// Method 2
	fmt.Println("=== Method 2 ===")
	Input = []int{0, 1, 1, 2, 1, 3}
	fmt.Println("Output --> ", HashTable(Input))
	fmt.Println("=== Method 2 ===")

	fmt.Println("")

	// Method 4
	fmt.Println("=== Method 4 ===")
	Input = []int{0, 1, 1, 2, 1, 3}
	fmt.Println("Output --> ", MarkTwoDiffNum(Input))
	fmt.Println("=== Method 4 ===")

	fmt.Println("")

}

// Author: HackerZ
// Time  : 2016/06/22 19:35

// Find One Number Appear More then Half.
// Input : [0, 1, 1, 2, 1, 3]
// Output: 1

// SortedAndFind Return the Middle Number in Sorted Array.
// 解法1：排序
func SortedAndFind(arrn []int) int {
	sort.Ints(arrn)
	return arrn[len(arrn)>>1]
}

// HashTable use HashTable.
// 解法2：散列表（哈希表）
func HashTable(arrn []int) int {
	Ht := make(map[int]int)
	for _, number := range arrn { // init hashtable
		Ht[number]++
	}

	// return The Number.
	var Max, target int
	for k, v := range Ht {
		if k > Max {
			target = v
		}
	}
	return target
}

// DelTwoDiffNum Delete Two Different Number even time.
// 解法3：每次删除两个不同的数
func DelTwoDiffNum(arrn []int) int {
	return arrn[0]
}

// MarkTwoDiffNum Mark Two Different Number.
// 解法4：记录两个数（解法3的优化版）
func MarkTwoDiffNum(arrn []int) int {
	candidate := arrn[0]
	nTimes := 1
	for i := 1; i < len(arrn) && nTimes <= len(arrn)>>1; i++ {
		/* when nTimes <= len(arrn)>>1
		   The candidate must be Which we Search.
		*/
		if nTimes == 0 {
			candidate = arrn[i]
			nTimes = 1
		} else {
			if candidate == arrn[i] {
				nTimes++
			} else {
				nTimes--
			}
		}
	}
	return candidate
}
