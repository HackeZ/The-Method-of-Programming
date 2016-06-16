package main

import "fmt"

func main() {
	arrn := []int{1, -2, 3, 10, -4, 7, 2, -5}
	// Method 1
	fmt.Println("=== Method 1 ===")
	fmt.Println("The Max Sub in Array --> ", MaxSubArray1(arrn, len(arrn)))
	fmt.Println("=== Method 1 ===")

	fmt.Println("")
	// Method 2
	fmt.Println("=== Method 2 ===")
	fmt.Println("The Max Sub in Array --> ", MaxSubArray2(arrn, len(arrn)))
	fmt.Println("=== Method 2 ===")

}

// Author: HackerZ
// Time  : 2016/06/16 14:15

// Max Sub Array
// Input : [1, -2, 3, 10, -4, 7, 2, -5]
// Output: 18

// MaxSubArray1 Brute force enumeration
// 解法1：蛮力枚举
func MaxSubArray1(arrn []int, n int) int {
	maxSum := arrn[0] // 全是负数的情况，返回最大的那个负数
	currSum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := i; k <= j; k++ {
				currSum += arrn[k]
			}
			if currSum >= maxSum {
				maxSum = currSum
			}

			// 清零，否则sum最终存放的是所有子数组的和
			currSum = 0
		}
	}
	return maxSum
}

// TIME: O(n^3)

// MaxSubArray2 Dynamic Programming
// 解法2：动态规划
func MaxSubArray2(arrn []int, n int) int {
	maxSum := arrn[0] // 全是负数的情况，返回最大的那个负数
	currSum := 0

	for j := 0; j < n; j++ {
		if currSum >= 0 {
			currSum += arrn[j]
		} else {
			currSum = arrn[j]
		}
		if currSum >= maxSum {
			maxSum = currSum
		}
	}
	return maxSum
}
