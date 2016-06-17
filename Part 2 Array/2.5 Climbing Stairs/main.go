package main

import "fmt"

func main() {
	// Method 1
	fmt.Println("=== Method 1 ===")
	fmt.Printf("Ways --> %d\n", Fibonacci(6))
	fmt.Println("=== Method 1 ===")
	fmt.Println("")
	// Method 2
	fmt.Println("=== Method 2 ===")
	fmt.Printf("Ways --> %d\n", ClimbStairs(6))
	fmt.Println("=== Method 2 ===")
}

// Author: HackerZ
// Time  : 2016/06/17 12:46

// Climbing Stairs
// There is n Stairs, You can climb it by 1 or 2 steps, How many Ways you can climb?
// Input:  6
// Output: 13

// Fibonacci
// 解法1：递归
func Fibonacci(n uint) int {
	result := []int{0, 1, 2}
	if n <= 2 {
		return result[n]
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// ClimbStairs
// 解法2：递推
func ClimbStairs(n uint) int {
	dp := [3]int{1, 1}
	if n < 2 {
		return 1
	}
	for i := 2; i <= int(n); i++ {
		dp[2] = dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = dp[2]
	}
	return dp[2]
}
