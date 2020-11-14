// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// type Duration int64 -> int64 자료형을 기반(underlying type)으로 한 커스텀 타입
// underlying type에서 사용 가능한 연산자, 표현, 사이즈가 동일하게 적용된다.
// 하지만 Duration과 int64를 형변환 없이 연산 또는 대입할 수는 없다, 완전히 다른 타입이기 때문

package main

import (
	"fmt"
	"time"
)

func main() {
	h, _ := time.ParseDuration("4h30m")

	// why would you want to create a new type?

	// 1- adding new methods to the type
	fmt.Println(h.Hours(), "hours")

	// 2- make it a distinct type for type-safety
	// you can't use the defined type
	//   with its underlying type directly.
	//
	// you need to convert one of them.
	var m int64 = 2
	h *= time.Duration(m)
	fmt.Println(h)

	// type of `h` and its underlying type are different
	fmt.Printf("Type of h: %T\n", h)
	fmt.Printf("Type of h's underlying type: %T\n", int64(h))
}
