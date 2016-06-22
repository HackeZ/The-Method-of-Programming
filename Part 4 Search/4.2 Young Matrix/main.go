package main

import (
	"fmt"
	"time"
)

func main() {
	Input := [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}
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
// func YoungMatrix1(arrn [][]int, target int) bool {
// 	fmt.Println("arrn ==> ", arrn)
// 	length := len(arrn[0])
// 	if length == 2 {
// 		for i := 0; i < length; i++ {
// 			for j := 0; j < length; j++ {
// 				if arrn[i][j] == target {
// 					return true
// 				}
// 			}
// 		}
// 		return false
// 	} else {
// 		for i := 0; i < length; i++ {
// 			j := i + 1
// 			if arrn[i][i] <= target && arrn[j][j] >= target {
// 				fmt.Println("arrn -->", arrn[i][i], arrn[j][j])
// 				fmt.Println("i --> ", i, " || j -->", j)
// 				fmt.Println("length -->", length)
// 				// return YoungMatrix1(arrn[j : length-1][j:length-1], target)
// 				show(arrn[0:1][0:2])
// 			}
// 		}
// 		return false
// 	}
// }

func show(arrn [][]int) {
	fmt.Println("show -->", arrn)
}

// YoungMatrix2 Location
// 解法2：定位法
func YoungMatrix2(arrn [][]int, searchKey int) bool {
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
