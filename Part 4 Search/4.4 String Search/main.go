package main

import "fmt"

func main() {
	// Method 1
	fmt.Println("=== Method 1 ===")
	strS, strP := "BBCABCDABABCDABDE", "ABCDABD"
	fmt.Println("Output --> ", ViolentMatch(strS, strP))
	fmt.Println("=== Method 1 ===")

	fmt.Println("")
}

// Author: HackerZ
// Time  :   2016/06/22 23:48

// String Search
// Input : "BBCABCDABABCDABDE" "ABCDABD"
// Output: 9

// ViolentMatch Violent Match.
// 解法1：蛮力匹配
func ViolentMatch(arrn, arrp string) int {
	sLen, pLen := len(arrn), len(arrp)
	var i, j int
	for i < sLen && j < pLen {
		if arrn[i] == arrp[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}
	if j == pLen {
		return i - j
	}
	return -1
}
