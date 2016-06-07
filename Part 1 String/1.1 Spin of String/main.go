package main

import "fmt"

func main() {
	// Method 1
	// -- star --
	s := "abcdef"
	m := 3
	for {
		if m > 0 {
			s = LeftShiftOne(s, len(s))
			m--
		} else {
			break
		}
	}
	fmt.Println("Method 1 --> ", s)
	// -- end --

	// Method 2
	// -- star --
	s = "abcdefgh"
	m = 4
	n := len(s)
	s = ReverseString(s, 0, m-1)
	s = ReverseString(s, m, n-1)
	s = ReverseString(s, 0, n-1)
	fmt.Println("Method 2 --> ", s)
	// -- end --
}

// Author: HackerZ
// Time:   2016/06/07 12:36

// Spin the String
// m: spin how many char to the end of string.
// Input:  abcdef
// Output: defabc

// LeftShiftOne Spin Char in each time.
// 解法1：蛮力移位法
func LeftShiftOne(s string, n int) string {
	sbytes := []byte(s)
	// save the first char
	t := sbytes[0]
	for i := 1; i < n; i++ {
		sbytes[i-1] = sbytes[i]
	}
	sbytes[n-1] = t
	fmt.Println("Method 1 --> ", string(sbytes))
	return string(sbytes)
}

// ReverseString Spin the String by 3 times.
// 解法2：三步反转
func ReverseString(s string, from, to int) string {
	sbytes := []byte(s)
	for {
		if from < to {
			sbytes[from], sbytes[to] = sbytes[to], sbytes[from]
			from++
			to--
		} else {
			break
		}
	}
	fmt.Println("Method 2 --> ", string(sbytes))
	return string(sbytes)
}
