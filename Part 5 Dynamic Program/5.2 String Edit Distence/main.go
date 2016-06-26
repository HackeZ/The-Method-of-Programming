package main

import "fmt"

var dp = make([][]int, 11, 11)

func main() {
	// Method 1
	fmt.Println("=== Method 1 ===")
	strS, strT := "ALGORITHM", "ALTRUISTIC"
	for i := 0; i < len(strS)+1; i++ {
		dp2 := make([]int, len(strT)+1) //可用循环对dp2赋值，默认建立初值为0
		dp[i] = dp2
	}
	fmt.Println(dp)
	fmt.Println("Output --> ", EditDistance(strS, strT))
	fmt.Println(dp)
	// dp[0][0] = 0

	fmt.Println("")
}

// Author: HackerZ
// Time  : 2016/06/26 15:32

/*String Edit Distence
* Given a Source and Target String.
* Handle Target String Method:
  - Insert a char in any location.
  - Replace and String.
  - Delete any String.

  By those Method, Using the least times to make Source eq Target.
*/

// Input  : ALGORITHM   ALTRUISTIC
// Output : 6

// EditDistence
// 解法1：动态规划
func EditDistance(strS, strT string) int {
	srcLength, targetLength := len(strS), len(strT)
	var i, j int
	for i = 1; i <= srcLength; i++ {
		dp[i][0] = i
	}
	for j = 1; j <= targetLength; j++ {
		dp[0][j] = j
	}

	for i = 1; i <= srcLength; i++ {
		for j = 1; j <= targetLength; j++ {
			if strS[i-1] == strT[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[srcLength][targetLength]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
