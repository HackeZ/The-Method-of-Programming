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
