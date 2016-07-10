package main

import "fmt"

func main() {
	// Method 1
	fmt.Println("Output --> ", IsInterleave("123", "456", "123456"))
}

// Author: HackerZ
// Time  : 2016-7-9 16:59

// Alternationg String.
/*
 * Input : s1 = "aabcc" s2 = "dbbca" s3 = "aadbbcbcac"
 *         s1 = "aabcc" s2 = "dbbca" s3 = "accabdbbca"
 *
 * Output: true
 *         false
 */

var dp = make([][]bool, 200, 200)

// IsInterleave Alternationg String or not.
func IsInterleave(s1, s2, s3 string) bool {
	n, m, s := len(s1), len(s2), len(s3)

	// In different length, Return False.
	if n+m != s {
		return false
	}

	// boolean [][]dp = new boolean[n + 1][m + 1];
	for i := 0; i < n+1; i++ {
		dp2 := make([]bool, m+1)
		dp[i] = dp2
	}

	// String can be null
	dp[0][0] = true

	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			if dp[i][j] || (i-1 >= 0 && dp[i-1][j] == true &&
				s1[i-1] == s3[i+j-1]) || (j-1 >= 0 && dp[i][j-1] == true &&
				s2[j-1] == s3[i+j-1]) {
				dp[i][j] = true
			} else {
				dp[i][j] = false
			}
		}
	}
	return dp[n][m]
}
