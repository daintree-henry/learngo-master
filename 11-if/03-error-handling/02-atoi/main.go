// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Atoi returns an error value
	// So, you should always check it
	// 인트로 변환하는 함수이기 때문에 문자를 넣으면 에러 발생한다

	// go run main.go 42

	n, err := strconv.Atoi(os.Args[1])

	fmt.Println("Converted number    :", n)
	fmt.Println("Returned error value:", err)
}
