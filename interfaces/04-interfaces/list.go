// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

type stuff interface {
	print()
} //-> 조건에 만족한 타입만 stuff로 다루어질 수 있다.

type list []stuff // print()함수를 가진 타입은 모두 사용 가능하다

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We're waiting for delivery 🚚.")
		return
	}

	for _, it := range l {
		// fmt.Printf("(%-10T) --> ", it)

		it.print()

		// you cannot access to the discount method of the game type.
		// `it` is a printer not a game.
		// it.discount(.5)
	}
}

// PREVIOUS CODE:

// type list []*game

// func (l list) print() {
// 	if len(l) == 0 {
// 		fmt.Println("Sorry. Our store is closed. We're waiting for the delivery 🚚.")
// 		return
// 	}

// 	for _, it := range l {
// 		it.print()
// 	}
// }
