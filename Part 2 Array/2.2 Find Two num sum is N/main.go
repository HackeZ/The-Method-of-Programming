package main

import (
	"fmt"
	"sort"
)

func main() {
	arrn := []int{1, 2, 4, 5, 7, 11, 15}
	n := 15
	num1, num2 := TwoSum3(arrn, n)
	fmt.Println(num1, " + ", num2)
}

// Author: HackerZ
// Time: 2016/06/13 17:23

// Find Two number which Sum is N.
// Input: [1, 2, 4, 5, 7, 11, 15]  15
// Output: 4 + 11

// Warning: Output only one possibility

// TwoSum1 Hash Map
// 解法1：散列映射
func TwoSum1(arrn []int, n int) (num1, num2 int) {
	var numMap map[int]int
	numMap = make(map[int]int)
	// init Hash Map
	for i := 0; i < len(arrn); i++ {
		result := n - arrn[i]
		numMap[result] = arrn[i]
	}

	// To find whether there are corresponding array keys
	for i := 0; i < len(arrn); i++ {
		if v, ok := numMap[arrn[i]]; ok {
			return arrn[i], v
		}
	}
	// Find Nothing
	return -1, -1
}

// TwoSum2 Sort And Clamp Approximation
// 解法2：排序夹逼 (Golang无法像C那样使用指针)
// void TwoSum(int data[], unsigned int length, int sum)
// {
//     //sort(s, s+n);   如果数组非有序的，那就事先排好序O(N log N)

//     int begin = 0;
//     int end = length - 1;

//     //俩头夹逼，或称两个指针两端扫描法，很经典的方法，O(N)
//     while (begin < end)
//     {
//         long currSum = data[begin] + data[end];

//         if (currSum == sum)
//         {
//             //题目只要求输出满足条件的任意一对即可
//             printf("%d %d\n", data[begin], data[end]);

//             //如果需要所有满足条件的数组对，则需要加上下面两条语句：
//             //begin++
//             //end--
//             break;
//         }
//         else{
//             if (currSum < sum)
//                 begin++;
//             else
//                 end--;
//         }
//     }
// }

// TwoSum3 Mean value judgment method
// 解法3：中值判定法 —— by HackerZ
func TwoSum3(arrn []int, n int) (num1, num2 int) {
	mid := n / 2
	sort.Ints(arrn)
	for i := 0; i < len(arrn); i++ {
		if mid < arrn[i] {
			if i < len(arrn)-i {
				num1, num2 = midNum(arrn[0:i], arrn[i:len(arrn)], n)
			} else {
				num1, num2 = midNum(arrn[i:len(arrn)], arrn[0:i], n)
			}
			return num1, num2
		}
	}
	return -1, -1
}

func midNum(arrS, arrL []int, n int) (num1, num2 int) {
	for i := 0; i < len(arrS); i++ {
		result := n - arrS[i]
		for num := range arrL {
			if result == num {
				return arrS[i], result
			}
		}
	}
	return -1, -1
}
