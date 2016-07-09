package main

import "fmt"

func main() {
	// Method 1
	fmt.Println("=== Method 1 ===")
	Input := [N][N]int{{0, 0, 3, 0, 2, 0}, {0, 0, 3, 0, 0, 0}, {0, 0, 3, 0, 0, 0}, {0, 0, 0, 0, 4, 0}, {0, 0, 0, 0, 4, 0}, {0, 0, 3, 0, 0, 0}}
	Output := MinPathSum(Input, len(Input[0]))
	fmt.Println("Output -->", Output)
	fmt.Println("=== Method 1 ===")
}

// Author: HackerZ
// Time  : 2016/06/27 19:47

// Max Number of Squares
/*
 * Input:   +----------------+
 *          |0  0  3  0  2  0|
 *          |0  0  3  0  0  0|
 *          |0  0  3  0  0  0|
 *          |0  0  0  0  4  0|
 *          |0  0  0  0  4  0|
 *          |0  0  3  0  0  0|
 *          +----------------+
 * Output:  22
 */

const (
	N   = 202        // N : The Largest of Squares
	inf = 1000000000 // inf : ∞
)

var dp [N * 2][N][N]int

// IsValid : Is Status Valid
func IsValid(step, x1, x2, n int) bool {
	y1, y2 := step-x1, step-x2
	return ((x1 >= 0) && (x1 < n) && (x2 >= 0) && (x2 < n) && (y1 >= 0) && (y1 < n) && (y2 >= 0) && (y2 < n))
}

// GetValue : Return Value of Dp.
//            Return -∞ When Out of Array.
func GetValue(step, x1, x2, n int) int {
	if IsValid(step, x1, x2, n) {
		return dp[step][x1][x2]
	}
	return (-inf)
}

// MinPathSum :
func MinPathSum(a [N][N]int, n int) int {
	P := n*2 - 2 // The End of Step
	var i, j, step int

	for i = 0; i < n; i++ {
		for j = i; j < n; j++ {
			dp[0][i][j] = -inf
		}
	}
	dp[0][0][0] = a[0][0]

	for step = 1; step <= P; step++ {
		for i = 0; i < n; i++ {
			for j = i; j < n; j++ {
				dp[step][i][j] = -inf
				if !IsValid(step, i, j, n) { // Not Valid
					continue
				}
				// DFS for Valid Location.
				if i != j {
					dp[step][i][j] = max(dp[step][i][j], GetValue(step-1, i-1, j-1, n))
					dp[step][i][j] = max(dp[step][i][j], GetValue(step-1, i-1, j, n))
					dp[step][i][j] = max(dp[step][i][j], GetValue(step-1, i, j-1, n))
					dp[step][i][j] = max(dp[step][i][j], GetValue(step-1, i, j, n))
					dp[step][i][j] += a[i][step-i] + a[j][step-j] // In Different location, plus 2 number.
				} else {
					dp[step][i][j] = max(dp[step][i][j], GetValue(step-1, i-1, j-1, n))
					dp[step][i][j] = max(dp[step][i][j], GetValue(step-1, i-1, j, n))
					dp[step][i][j] = max(dp[step][i][j], GetValue(step-1, i, j-1, n))
					dp[step][i][j] += a[i][step-i] // In Same location, plus 1 number.
				}
			}
		}
	}
	return dp[P][n-1][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
