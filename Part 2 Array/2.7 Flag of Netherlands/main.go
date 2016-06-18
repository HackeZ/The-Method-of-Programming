package main

import "fmt"

func main() {
	fmt.Println("=== Method 1 ===")
	Input := []int{2, 0, 1, 0, 0, 2, 1, 1, 2}
	fmt.Println("Input --> ", Input)
	QuickSort(Input, len(Input))
	fmt.Println("Output --> ", Input)
	fmt.Println("=== Method 1 ===")
}

// Author: HackerZ
// Time  : 2016/06/17 14:16

// Flag of Netherlands
// Means that   R - 0
//              W - 1
//              B - 2
// Sort the RWB

// Input : [2, 0, 1, 0, 0, 2, 1, 1, 2]
// Output: [0, 0, 0, 1, 1, 1, 2, 2, 2]

// QuickSort Sort the RWB
// 解法1：快速排序
func QuickSort(arrn []int, length int) {
	begin, current, end := 0, 0, length-1
	for {
		if current > end {
			break
		}
		if arrn[current] == 0 {
			arrn[begin], arrn[current] = arrn[current], arrn[begin]
			current++
			begin++
		} else if arrn[current] == 1 {
			current++
		} else if arrn[current] == 2 {
			arrn[end], arrn[current] = arrn[current], arrn[end]
			end--
		}
	}
}
