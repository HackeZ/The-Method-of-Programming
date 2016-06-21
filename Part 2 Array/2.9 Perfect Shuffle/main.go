package main

import "fmt"

func main() {

	// Method 1
	fmt.Println("=== Method 1 ===")
	Input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	Output := Shuffle1(Input, len(Input)/2)
	fmt.Println("Output --> ", Output)
	fmt.Println("=== Method 1 ===")

	fmt.Println("")

	// Method 2
	fmt.Println("=== Method 2 ===")
	Input = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	Output = Shuffle2(Input, len(Input)/2)
	fmt.Println("Output --> ", Output)
	fmt.Println("=== Method 2 ===")

	fmt.Println("")

	// Method 3
	fmt.Println("=== Method 3 ===")
	Input = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	Output = LocationReplace(Input, len(Input)/2-1)
	fmt.Println("Output --> ", Output)
	fmt.Println("=== Method 3 ===")

	fmt.Println("")

	// Method 4
	fmt.Println("=== Method 4 ===")
	Input = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	Output = Shuffle(Input, len(Input)/2)
	fmt.Println("Output --> ", Output)
	fmt.Println("=== Method 4 ===")

}

// Author: HackerZ
// Time  : 2016/06/20 01:14

// Perfect Shuffle Means that
//      {a1, a2, ..., an, b1, b2, ..., bn} ==> {a1, b1, a2, b2, ..., an, bn}
// Tips:   TIME:O(n) && SPACE:O(1)

// Input : {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// Output: {0, 4, 1, 5, 2, 6, 3, 7}

// Shuffle1 Brute force method(One Step After One Step)
// 解法1：蛮力解法（步步为营）
func Shuffle1(arrn []int, n int) []int {
	var i int
	for {
		if n-i == 1 {
			break
		} else {
			fmt.Println("n +", i, "(", arrn[n+i], ")", "move forawrd", n-i-1, "steps.")
			t := arrn[n+i]
			for k := 1; k <= n-i-1; k++ {
				arrn[n+i-k+1] = arrn[n+i-k]
			}
			arrn[2*i+1] = t
			i++
		}
	}
	return arrn
}

//TIME:O(n^2)

// Shuffle2 Brute force method(Middle Swap)
// 解法2：蛮力解法（中间交换）
func Shuffle2(arrn []int, n int) []int {
	for i := 1; i <= n-1; i++ {
		point := 0
		fmt.Println("The", i, "Time to Swap.")
		for k := 0; k < i; k++ {
			arrn[n-i+point+1], arrn[n-i+point] = arrn[n-i+point], arrn[n-i+point+1]
			fmt.Println(arrn[n-i+point+1], "Swap", arrn[n-i+point])
			point += 2
		}
	}
	return arrn
}

// TIME:O(n^2)

// LocationReplace by Peiyush Jain
// 解法3：位置置换
func LocationReplace(arrn []int, n int) []int {
	n2 := n * 2
	var i int
	var b = make([]int, 1000)
	for i = 1; i <= n2; i++ {
		// fmt.Println("b --> ", (i*2)%(n2+1))
		b[(i*2)%(n2+1)] = arrn[i]
	}
	for i = 1; i <= n2; i++ {
		arrn[i] = b[i]
	}
	return arrn
}

// TIME:O(n)  SPACE:O(n)

// Shuffle perface shuffle
// 解法4：完美洗牌算法
// func Shuffle(arrn []int, n int) []int {
// 	n2 := n * 2
// 	arrn = PerfaceShuffle(arrn, n)
// 	for i := 2; i <= n2; i += 2 {
// 		arrn[i-1], arrn[i] = arrn[i], arrn[i-1]
// 	}
// 	return arrn
// }

// func PerfaceShuffle(arrn []int, n int) []int {
// 	var n2, m, i, k, t int
// 	for n > 1 {
// 		// First Step.
// 		n2 = n * 2
// 		for k, m = 0, 1; n2/m >= 3; {
// 			k++
// 			m *= 3
// 		}
// 		m /= 2
// 		// 2m = 3^k - 1 , 3^k <= 2n < 3^(k+1)

// 		// Second Step.
// 		RightRotate(arrn[m:], m, n)

// 		// Third Step.
// 		for i, t = 0, 3; i < k; {
// 			CycleLeader(arrn, t, m*2+1)
// 			i++
// 			t *= 3
// 		}

// 		// Forth Step.
// 		arrn = arrn[m*2:]
// 		n -= m
// 	}
// 	// n = 1
// 	arrn[1], arrn[2] = arrn[2], arrn[1]
// 	return arrn
// }

// // reverse the []int. TIME:O(to-from)
// func reverse(arrn []int, form, to int) {
// 	for form < to {
// 		arrn[form], arrn[to] = arrn[to], arrn[form]
// 		form++
// 		to--
// 	}
// }

// // RightRotate rotate right num steps.TIME:O(n)
// func RightRotate(arrn []int, num, n int) {
// 	reverse(arrn, 1, n-num)
// 	reverse(arrn, n-num+1, n)
// 	reverse(arrn, 1, n)
// }

// // CycleLeader Cycle Algorithm.
// func CycleLeader(arrn []int, from, mod int) []int {
// 	for i := from * 2 % mod; i != from; i = i * 2 % mod {
// 		arrn[from], arrn[i] = arrn[i], arrn[from]
// 	}
// 	return arrn
// }
