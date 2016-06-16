package main

import (
	"container/list"
	"fmt"
)

var list1 = list.New()

func main() {
	// Method 1
	fmt.Println("=== Method 1 ===")
	SumofkNumber1(15, 8)
	fmt.Println("=== Method 1 ===")
}

// Author: HackerZ
// Time  : 2016/06/15 19:15

// Sum of K Number
// Input : 	15  8 // sum n
// Output: 	8 + 7
//			6 + 8 + 1
//			5 + 8 + 2
//			4 + 8 + 3
//			2 + 4 + 8 + 1
//			6 + 7 + 2
//			5 + 7 + 3
//			2 + 5 + 7 + 1
//			3 + 4 + 7 + 1
//			5 + 6 + 4
//			3 + 5 + 6 + 1
//			3 + 4 + 6 + 2
//			2 + 3 + 4 + 5 + 1

// SumofkNumber1 n Problem turn to n-1 Problem
// 解法1：n问题转换为n-1问题
func SumofkNumber1(sum, n int) {
	// 递归出口
	if n <= 0 || sum <= 0 {
		return
	}
	// 输出找到的结果
	if sum == n {
		for iter := list1.Front(); iter != nil; iter = iter.Next() {
			fmt.Printf("%v + ", iter.Value)
		}
		fmt.Printf("%d\n", n)
	}

	list1.PushFront(n)        // 典型的 01 背包问题
	SumofkNumber1(sum-n, n-1) // 放n，前n-1个数填满sum-n
	list1.Remove(list1.Front())
	SumofkNumber1(sum, n-1) // 不放n，前n-1个数填满sum
}

// SumofkNumber2 Backtrackruning
// 解法2：回溯剪枝
// func SumofkNumber2()
