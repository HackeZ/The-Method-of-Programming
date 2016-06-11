package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	// strconv.Atoi("123")
	n, err := StrToInt("-8 19")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("StrToInt --> ", n)
	}
}

// Author: HackerZ
// Time: 2016/06/10

// StrToInt
// Input: "123" ==> string
// Output: 123  ==> int

// StrToInt foce on detail
func StrToInt(str string) (int, error) {
	// is nil
	if len(str) == 0 {
		return 0, errors.New("Error: Type of var.")
	}

	// Handle Space
	str = strings.Replace(str, " ", "", -1)
	// Handle +/-
	sign := 1
	if str[0] == '+' || str[0] == '-' {
		if str[0] == '-' {
			sign = -1
		}
		str = str[1:]
	}

	n := 0
	// Must be a number
	for strbyte := range str {
		if str[strbyte] < 48 || str[strbyte] > 57 {
			return 0, errors.New("Error: Not Number.")
		}
		n = n*10 + int(str[strbyte]-'0')
	}

	return n * sign, nil
}
