// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// normal declaration use cases

// -----------------------------------------------------
// when you need a package scoped variable
// 패키지 스코프의 변수가 필요한 경우
// -----------------------------------------------------

// version := 0 // YOU CAN'T
var version int

func main() {

	// -----------------------------------------------------
	// if you don't know the initial value
	// 초기값이 정해지지 않은 경우
	// -----------------------------------------------------

	// DON'T DO THIS:
	// score := 0

	// DO THIS:
	// var score int

	// -----------------------------------------------------
	// group variables for readability
	// 가독성을 위해 연관있는 변수는 한번에 선언 가능
	// -----------------------------------------------------

	// var (
	// 	video    string

	// 	duration int
	// 	current  int
	// )
}
