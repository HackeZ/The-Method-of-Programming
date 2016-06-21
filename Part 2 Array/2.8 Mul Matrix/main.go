package main

import "fmt"

func main() {
	// Method 1
	fmt.Println("=== Method 1 ===")
	matrixA := [][]int{{1, 2}, {3, 4}}
	matrixB := [][]int{{0, 1}, {0, 0}}
	matrixC := [][]int{{0, 0}, {0, 0}}
	MelMatrix1(matrixA, matrixB, matrixC, 2)
	fmt.Println("Output --> ", matrixC)
	fmt.Println("=== Method 1 ===")
}

// Author: HackerZ
// Time  : 2016/06/18 22:23

// Mul Matrix
// Input:   +----+      +----+
//          |1  2|      |0  1|
//          |3  4|      |0  0|
//          +----+      +----+

// Output:  +----+
//          |0  1|
//          |0  3|
//          +----+

// MelMatrix1 Brute force method (the same dimension n)
// 解法1：蛮力解法（相同维数n）
func MelMatrix1(matrixA, matrixB, matrixC [][]int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			matrixC[i][j] = 0
			for k := 0; k < n; k++ {
				matrixC[i][j] += matrixA[i][k] * matrixB[k][j]
			}
		}
	}
}

// MelMatrix2 Strassen Algorithm.
