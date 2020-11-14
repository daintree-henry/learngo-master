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
	var brand string

	// prints the string in quoted-form like this ""
	fmt.Printf("%q\n", brand)

	brand = "Google"
	fmt.Printf("%q\n", brand)
	// Printf의 동사 자리에 순서대로 변수가 대체된다
	// [verbs]
	// %q : 문장
	//
	// /n : 새로운 라인
}
