package main

import "fmt"

func main() {
	Input := []int{1, 4, 6, 7, 11, 16, 31, 55}
	Output := BinarySearch(Input, len(Input), 16)
	fmt.Println("Output --> ", Output)
}

// Author: HackerZ
// Time  : 2016/06/22 12:38

// Binary Search
// Warning: Must use Sorted Array.
// Input : [1, 4, 6, 7, 11, 16, 31, 55]  16
// Output: 5

// BinarySearch Must use Sorted Array.
// 解法：正确的二分查找
func BinarySearch(arrn []int, n, value int) int {
	left, right := 0, n-1
	// if right eq n. There are two place to change.
	// 1) for left < right {
	// 2) in for, when arrn[middle] > value, right = mid.
	for left <= right {
		middle := left + ((right - left) >> 1)
		if arrn[middle] > value {
			right = middle - 1 // change the right.
		} else if arrn[middle] < value {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}
