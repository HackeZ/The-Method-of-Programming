package main

import "fmt"

func main() {
	str := "12212321"
	len := LongestPalindrome1(str, len(str))

	fmt.Println(str, " Longest Palindrome --> ", len)
}

// Author: HackerZ
// Time: 2016/06/11 13:45

// Longest Palindrome
// Input:  "12212321"
// Output: 5

// LongestPalindrome1 sum Longest Palindrome
// 解法1：中心扩展法
func LongestPalindrome1(str string, n int) int {
	var max, c int
	if len(str) == 0 || n < 1 {
		return 0
	}

	// i: center of Palindrome
	for i := 0; i < n; i++ {
		// len of longest Palindrome is odd number
		for j := 0; (i-j >= 0) && (i+j < n); j++ {
			if str[i-j] != str[i+j] {
				break
			}
			c = j*2 + 1
		}
		if c > max {
			max = c
		}

		// len if longest Palindrome is even
		for j := 0; (i-j >= 0) && (i+j+1 < n); j++ {
			if str[i-j] != str[i+j+1] {
				break
			}
			c = j*2 + 2
		}
		if c > max {
			max = c
		}
	}
	return max
}

// Manacher sum longest Palindrome
// 解法2：Manacher算法
