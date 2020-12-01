// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

// file scope - 파일 내에서만 유효하다
import "fmt"

// package scope - 같은 패키지 내에서 유효
const ok = true

// package scope  - 같은 패키지 내에서 유효
func main() { // block scope starts - 해당 함수 내에서만 유효

	var hello = "Hello"

	// hello and ok are visible here
	fmt.Println(hello, ok)

} // block scope ends
