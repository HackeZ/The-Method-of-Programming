package main

import "fmt"

const str = "abc"

func main() {
	fmt.Println("=== Method 1 ===")
	CalcAllPermutation1(str, 0, len(str)-1)
	fmt.Println("=== Method 1 ===")
}

// Authot: HackerZ
// Time: 2016/06/09 11:47

// String Permutation
// Input: 'abc'
// Output: ['abc', 'acb', 'bac', 'bca', 'cab', 'cba']

// CalcAllPermutation1 String Permutation
// 解法1：递归实现
func CalcAllPermutation1(perm string, from, to int) {
	if to < 1 {
		return
	}
	if from == to {
		for i := 0; i <= to; i++ {
			fmt.Printf("%c", perm[i])
		}
		fmt.Printf("\n")
	} else {
		permBytes := []byte(perm)
		for j := from; j <= to; j++ {
			permBytes[j], permBytes[from] = permBytes[from], permBytes[j]
			perm = string(permBytes)
			CalcAllPermutation1(perm, from+1, to)
			permBytes := []byte(perm)
			permBytes[j], permBytes[from] = permBytes[from], permBytes[j]
			perm = string(permBytes)
		}
	}
}

// CalcAllPermutation2 String Permutation
// 解法2：字典序排序
