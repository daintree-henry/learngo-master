// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	const min = 1

	max := 5 + min
	// above line equals to this:
	// max := int(5) + int(min)
	// 참고할 자료형이 없기 때문에 기본형인 int로 할당

	fmt.Printf("Type of max: %T\n", max)
}
