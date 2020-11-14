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
	const (
		//상수형을 위한 자동 증가값 생성기
		monday = iota
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday,
		saturday, sunday)
}
