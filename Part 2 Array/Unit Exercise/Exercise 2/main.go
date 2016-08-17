package main

import "fmt"

func main() {
	// Method 1
	fmt.Println("=== Method 1 ===")

	fmt.Println("Output --> ")
	fmt.Println("=== Method 1 ===")

}

// Author : HackerZ
// Time   : 2016/08/17 15:06

// Find Three Numbers that Appear Only Once.
// In Array Only Three Numbers Appear Once but Other Twice. Find This Three Numbers.

// Input : {1, 2, 3, 4, 4, 5, 2}
// Output : 1  3  5

// FindThreeNumbersOnce
// @param  arrInt []int
// @return elem int
func FindThreeNumbersOnce(arrInt []int) (elem int) {
	return arrInt[0]
}
