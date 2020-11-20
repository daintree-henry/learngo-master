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
)

func main() {
	city := os.Args[1]

	switch city {
	case "Paris":
		fmt.Println("France")
		// break // unnecessary in Go

		// 케이스 내부에서 선언한 변수는 하나의 케이스 안에서만 사용 가능하다
		// vip := true
		// fmt.Println("VIP trip?", vip)

	case "Tokyo":
		fmt.Println("Japan")
		// fmt.Println("VIP trip?", vip)
	}

	// if city == "Paris" {
	// 	fmt.Println("France")
	// } else if city == "Tokyo" {
	// 	fmt.Println("Japan")
	// }
}
