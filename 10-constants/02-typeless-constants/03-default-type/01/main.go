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
	const min int32 = 1

	max := 5 + min
	// above line equals to this:
	// max := int32(5) + min
	// 변수 할당 시 이미 자료형이 정의된 부분이 있는지 찾아서 자동으로 변환
	// 아무 정보가 없는 경우 기본 자료형으로 할당

	fmt.Printf("Type of max: %T\n", max)
}
