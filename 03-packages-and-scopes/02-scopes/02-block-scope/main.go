// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

func nope() { // block scope starts

	// hello and ok are only visible here
	const ok = true
	var hello = "Hello"

	_ = hello
} // block scope ends

func main() { // block scope starts

	// hello and ok are not visible here

	// ERROR: 스코프가 유효하지 않은 변수를 참조하여 에러 발생
	// fmt.Println(hello, ok)

} // block scope ends
