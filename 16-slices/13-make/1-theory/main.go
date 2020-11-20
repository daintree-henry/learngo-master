// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	prettyslice "github.com/inancgumus/prettyslice"
)

func main() {
	prettyslice.PrintBacking = true

	prettyslice.Show("make([]int, 3)", make([]int, 3))
	prettyslice.Show("make([]int, 3, 5)", make([]int, 3, 5))
	//make함수를 사용해 슬라이스 생성 시 cap을 지정할 수 있다

	s := make([]int, 0, 5)
	prettyslice.Show("make([]int, 0, 5)", s)

	s = append(s, 42)
	prettyslice.Show("s = append(s, 42)", s)
}
