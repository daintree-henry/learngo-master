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
	// = is the assignment operator
	// when used within a variable declaration, it
	// initializes the variable to the given value

	// here, Go initializes the safe variable to true

	// OPTION #1 (option #2 is better)
	// var safe bool = true

	// OPTION #2
	var safe = true
	// Go 가 자동으로 bool 자료형으로 초기화한다.
	// => safe := true 형태로 축약하여 사용한다

	fmt.Println(safe)
}
