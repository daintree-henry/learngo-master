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

	s "github.com/inancgumus/prettyslice"
)

func main() {
	// golang.org/src/runtime/slice.go

	// 슬라이스는 Pointer(첫 번째 요소의 메모리 주소), Length(배열의 길이), Capacity <- backing array의 끝값까지의 거리, cap(slice)로 조회 가능, cap을 넘는 인덱스에 접근 불가능하다 의 값을 가진다.
	// 초기화되지 않은 nil 슬라이스도 슬라이스 헤더를 가진다, 3개의 값이 모두 0이다.

	// each int element is 4 bytes (on 64-bit)
	//
	// let's say the ages point to 1000th.
	// ages[1:] will point to 1004th
	// ages[2:] will point to 1008th and so on.
	//
	// they all will be looking at the same
	// backing array.
	//

	ages := []int{35, 15, 25}
	red, green := ages[0:1], ages[1:3]

	s.Show("ages", ages)
	s.Show("red", red)
	s.Show("green", green)

	fmt.Println(red[0])
	// fmt.Println(red[1]) // error
	// fmt.Println(red[2]) // error

	{
		var ages []int
		s.Show("nil slice", ages)

		// or just:
		s.Show("nil slice", []int(nil))
	}
}
