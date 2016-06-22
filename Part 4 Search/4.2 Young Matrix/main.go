package main

import (
	"fmt"
	"time"
)

func main() {
	Input := [4][4]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}
	// for i := 0; i != len(Input[6]); i++ {
	// 	fmt.Println("Y")
	// }
	Output := YoungMatrix2(Input, 6)
	fmt.Println("Output --> ", Output)
}

// Author: HackerZ
// Time  : 2016/06/22 12:19

// Young Matrix
// Input:   +----------+    6
//          |1  2  8  9|
//          |2  4  9 12|
//          |4  7 10 13|
//          |6  8 11 15|
//          +----------+

// Output: true

// YoungMatrix1 Divide and Conquer.
// 解法1：分治法
func YoungMatrix1(arrn [][]int, target int) bool {
	if len(arrn[0]) == 2 {
		for i := 0; i < len(arrn[0]); i++ {
			for j := 0; j < len(arrn[0]); j++ {
				if arrn[i][j] == target {
					return true
				}
			}
		}
		return false
	} else {
		for i := 0; i < len(arrn[0]); i++ {
			j := i + 1
			if arrn[i][i] <= target && arrn[j][j] >= target {
				return YoungMatrix1(arrn[j:len(arrn[0])][j:len(arrn[0])], target)
			}
		}
		return false
	}
}

// YoungMatrix2 Location
// 解法2：定位法
func YoungMatrix2(arrn [4][4]int, searchKey int) bool {
	i, j := 0, len(arrn[0])-1
	var1 := arrn[i][j]
	for {
		fmt.Println("var1 --> ", var1)
		if var1 == searchKey {
			return true
		} else if var1 < searchKey && i < len(arrn[0])-1 {
			i++
			var1 = arrn[i][j]
		} else if var1 > searchKey && j > 0 {
			j--
			var1 = arrn[i][j]
		} else {
			return false
		}
		time.Sleep(time.Millisecond * 500)
	}
}
