// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	s "github.com/inancgumus/prettyslice"
)

func main() {
	s.PrintBacking = true

	ages := []int{35, 15}
	s.Show("ages", ages)
	// len=2, cap=2, ptr

	//append 함수는 새로운 길이만큼의 backing array로 재할당한다 <- 재할당 횟수를 줄이기 위해
	//새로운 cap 자리도 *2만큼 비워둔다. 값은 커질수록 배수가 작아진다
	//len=3,cap=4,ptr=새로운포인터
	ages = append(ages, 5)
	s.Show("append(ages, 5)", ages)
}
