package main

import "fmt"

func main() {
	// Method 1
	Input := []int{1, 2, 6, 9, 7}
	fmt.Println("=== Method 1 ===")
	Output := OddEvenSort1(Input, uint(len(Input)))
	fmt.Println("Output --> ", Output)
	fmt.Println("=== Method 1 ===")

	fmt.Println("")

	// Method 2
	Input = []int{1, 2, 6, 9, 7}
	fmt.Println("=== Method 2 ===")
	OddEvenSort2(Input, 0, len(Input)-1)
	fmt.Println("Output --> ", Input)
	fmt.Println("=== Method 2 ===")
}

// Author: HackerZ
// Time  : 2016/06/17 11:04

// Odd Even Sort
// Tips: Even Before Odd && TIME eq O(n)
// Input : {1, 2, 6, 9, 7}
// Output: {1, 9, 7, 2, 6}  ||  {1, 7, 9, 6, 2}

// isOddNumber return true if is Odd Number.
func isOddNumber(data int) bool {
	return (data & 1) == 1
}

// OddEvenSort1 Head and Tail Point sort To the middle of the scan.
// 解法1：一头一尾指针往中间扫描
func OddEvenSort1(arrn []int, length uint) []int {
	if len(arrn) == 0 || length == 0 {
		return []int{}
	}
	head, tail := 0, len(arrn)-1
	for {
		if head > tail {
			break
		}
		if isOddNumber(arrn[head]) {
			// 如果head指向的是奇数，正常，向右移
			head++
		} else if !isOddNumber(arrn[tail]) {
			// 如果tail指向的是偶数，正常，向左移
			tail--
		} else {
			arrn[head], arrn[tail] = arrn[tail], arrn[head]
		}
	}
	return arrn
}

// TIME:O(n)

// OddEvenSort2 Front and Back Point sort To the End of the scan.
// 解法2：一前一后指针往后扫描
func OddEvenSort2(arrn []int, lo, hi int) {
	i := lo - 1
	for j := lo; j < hi; j++ {
		if isOddNumber(arrn[j]) {
			i++
			arrn[i], arrn[j] = arrn[j], arrn[i]
		}
	}
	arrn[i+1], arrn[hi] = arrn[hi], arrn[i+1]
}
