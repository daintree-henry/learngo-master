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
	// constants repeat the previous type and expression
	const (
		min int = 1
		max     // int = 1 생략 시 자동으로 이전 오퍼레이터를 반복한다
	)
	// 해당 속성은 iota에서 활용됨

	fmt.Println(min, max)
	// 1,1

	// print the types of min and max
	fmt.Printf("%T %T\n", min, max)
}
