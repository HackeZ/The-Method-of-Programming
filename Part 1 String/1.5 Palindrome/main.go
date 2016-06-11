package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "我爱我"
	ok := IsPalindrome(str)
	if ok {
		fmt.Println("is Palindrome.")
	} else {
		fmt.Println("not Palindrome.")
	}
}

// Author: HackerZ
// Time: 2016/06/10 23:23

// isPalindrome
// Input: madam  ||  我爱我  ||  StrToInt
// Output: true  ||  true    ||  false

// Warning: In Golang, Because of the string point can not use like C, so we use recursion to solve the Problem

// IsPalindrome in C
// 解法1：两头往中间扫
// bool IsPalindrome(const char *s, int n) {
// 	// illegal input
// 	if(s == NULL || n < 1) {
// 		return false;
// 	}
// 	const char* front, *back;

// 	// init the front and back point
// 	front = s;
// 	back = s + n - 1;

// 	while(front < back) {
// 		if *front != *back {
// 			return false;
// 		}
// 		++front;
// 		++back;
// 	}
// 	return true;
// }
// TIME: O(n)  SPACE: O(1)

// IsPalindrome in C
// 解法2：中间往两头扫
// bool IsPalindrime(const char *s, int n) {
// 	if (s == NULL || n < 1) {
// 		return false;
// 	}
// 	const char* first, *second;

// 	// m定位到字符串的中间位置
// 	int m = ((n >> 1) - 1) >= 0 ? (n >> 1) - 1 : 0;
// 	first = s + m;
// 	second = s + n - 1 - m;

// 	while (first >= s) {
// 		if (*first != *second) {
// 			return false;
// 		}
// 		--first;
// 		++second;
// 	}
// 	return true;
// }
// TIME: O(n)  SPACE: O(1)

// IsPalindrime in Golang
// 解法3：递归 from http://my.oschina.net/guonaihong/blog/419883
func doPalindrome(s string) bool {
	if utf8.RuneCountInString(s) <= 1 {
		return true
	}

	word := strings.Trim(s, "\t \r\n\v")
	first, sizeOfFirst := utf8.DecodeRuneInString(word)
	last, sizeOfLast := utf8.DecodeLastRuneInString(word)

	if first != last {
		return false
	}
	return doPalindrome(word[sizeOfFirst : len(word)-sizeOfLast])
}

func IsPalindrome(word string) bool {
	s := ""
	s = strings.Trim(word, "\t \r\n\v")
	if len(s) == 0 || len(s) == 1 {
		return false
	}
	return doPalindrome(s)
}
