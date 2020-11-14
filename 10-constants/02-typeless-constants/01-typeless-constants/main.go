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
	// min and pi are typeless constants
	const min = 1 + 1
	const pi = 3.14 * min
	// 변수일 때는 float * int 로 에러가 발생하지만 해당 계산은 가능
	// 타입 없이 선언한 상수는 untyped(typeless) 상수가 된다.

	fmt.Println(min, pi)
}
