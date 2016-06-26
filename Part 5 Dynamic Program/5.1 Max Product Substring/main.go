package main

import (
	"fmt"
)

func main() {
	// Method 1
	fmt.Println("=== Method 1 ===")
	Input := []float64{-2.5, 4, 0, 3, 0.5, 8, -1}
	fmt.Println("Output --> ", ViolentRoundRobin(Input, len(Input)))
	fmt.Println("=== Method 1 ===")

	fmt.Println("")

	// Method 2
	fmt.Println("=== Method 2 ===")
	Input = []float64{-2.5, 4, 0, 3, 0.5, 8, -1}
	fmt.Println("Output --> ", MaxProductSubstring(Input, len(Input)))
	fmt.Println("=== Method 2 ===")
}

// Author: HackerZ
// Time  : 2016/06/26 13:41

// Max Product Substring
// Input : [-2.5, 4, 0, 3, 0.5, 8, -1]
// Output: 12
// Tips  --> 12 eq 3 * 0.5 * 8

// ViolentRoundRobin Round Robin Array
// 解法1：蛮力轮询
func ViolentRoundRobin(arrn []float64, length int) float64 {
	maxResult := arrn[0]
	for i := 0; i < length; i++ {
		var x float64 = 1
		for j := i; j < length; j++ {
			x *= arrn[j]
			if x > maxResult {
				maxResult = x
			}
		}
	}
	return maxResult
}

// TIME:O(n^2)

// MaxProductSubstring Max Product Substring
// 解法2：动态规划
func MaxProductSubstring(arrn []float64, length int) float64 {
	maxEnd, minEnd, maxResult := arrn[0], arrn[0], arrn[0]
	for i := 1; i < length; i++ {
		end1, end2 := maxEnd*arrn[i], minEnd*arrn[i]
		maxEnd = max(max(end1, end2), arrn[i])
		minEnd = min(min(end1, end2), arrn[i])
		maxResult = max(maxResult, maxEnd)
	}
	return maxResult
}

// TIME:O(n)

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
