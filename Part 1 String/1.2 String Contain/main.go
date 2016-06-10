package main

import (
	"fmt"
	"sort"
)

// == SortString ==
type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

// SortString sort the string by package sort
func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

// == SortString ==

func main() {
	a, b := "ABCD", "BAD"
	// Method 1
	if StringContain1(a, b) {
		fmt.Println("Method 1 --> True")
	} else {
		fmt.Println("Method 1 --> False")
	}

	// Method 2
	if StringContain2(a, b) {
		fmt.Println("Method 2 --> True")
	} else {
		fmt.Println("Method 2 --> False")
	}
}

// Author: HackerZ
// Time  : 2016/06/07 12:06

// String Contain
// Input : a -> "ABCD" b -> "BAD" || a -> "ABCD" b -> "BCE" || a -> "ABCD" b -> "AA"
// Output: true || false || true

// StringContain1 String Contain
// 解法1：蛮力轮询
func StringContain1(a, b string) bool {
	for i := 0; i < len(b); i++ {
		var j int
		for j = 0; j < len(a) && a[j] != b[i]; j++ {
		}
		if j >= len(a) {
			return false
		}
	}
	return true
}

// StringContain2 String Contain
// 解法2：排序后轮询
func StringContain2(a, b string) bool {
	a = SortString(a)
	b = SortString(b)
	for pa, pb := 0, 0; pb < len(b); {
		for {
			if pa < len(a) && a[pa] < b[pb] {
				pa++
			} else {
				break
			}
		}
		if pa >= len(a) || a[pa] > b[pb] {
			return false
		}
		pb++
	}
	return true
}

// StringContain3 String Contain
// 解法3：素数相乘（不建议，因为素数的乘积结果非常大）
func StringContain3(a, b string) bool {
	p := [26]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	f := 1
	for i := 0; i < len(a); i++ {
		x := p[a[i]-'A']
		// 排除相同的字符，提高效率
		if f%x == 1 {
			f *= x
		}
	}
	for i := 0; i < len(b); i++ {
		x := p[a[i]-'A']
		if f%x == 1 {
			return false
		}
	}
	return true
}

// StringContain4 String Contain
// 解法4：位运算法
func StringContain4(a, b string) bool {
	return true
}
