package main

import "fmt"

var next = make([]int, 7)

func main() {
	// Method 1
	fmt.Println("=== Method 1 ===")
	strS, strP := "BBC ABCDAB ABCDABDE", "ABCDABD"
	fmt.Println("Output --> ", ViolentMatch(strS, strP))
	fmt.Println("=== Method 1 ===")

	fmt.Println("")

	// Method 2
	fmt.Println("=== Method 2 ===")
	strS, strP = "BBC ABCDAB ABCDABDE", "ABCDABD"
	GetNextval(strP, next)
	fmt.Println("Output --> ", KmpSearch(strS, strP))
	fmt.Println("=== Method 2 ===")
}

// Author: HackerZ
// Time  :   2016/06/22 23:48

// String Search
// Input : "BBC ABCDAB ABCDABDE" "ABCDABD"
// Output: 11

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

// KmpSearch KMP Algorithm.
// 解法2：KMP算法
func KmpSearch(S, P string) int {
	i, j := 0, 0
	sLen, pLen := len(S), len(P)
	fmt.Println("next --> ", next)
	// time.Sleep(time.Minute)
	for i < sLen && j < pLen {
		/* 	If j = -1 or string match Success.
		//  i++
		//  j++
		*/
		if j == -1 || S[i] == P[j] {
			i++
			j++
		} else {
			/*	If j != -1 also string match Fail.
				i will not change and j = next[j]
			*/
			j = next[j]
		}
	}
	if j == pLen {
		return i - j
	}
	return -1
}

// GetNextval Get next Array in KMP.
func GetNextval(P string, next []int) {
	pLen := len(P)
	next[0] = -1
	k, j := -1, 0

	for j < pLen-1 {
		/*	P[j] is prefix,
			P[k] is suffix.
		*/
		if k == -1 || P[j] == P[k] {
			j++
			k++
			if P[j] != P[k] {
				next[j] = k
			} else {
				/*	Beause of P[j] = P[next[j]] cannot appear.
					Need to continue to recursive,
					k = next[k] = next[next[k]]
				*/
				next[j] = next[k]
			}
		} else {
			k = next[k]
		}
	}
}
