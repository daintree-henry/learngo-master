// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

type book struct {
	title string
	price float64
}

// book 타입에 속한 함수 선언
func (b book) print() {
	// b is a copy of the original `book` value here.
	fmt.Printf("%-15s: $%.2f\n", b.title, b.price)
}

// ----------------------------------------------------------------------------
// + you can use the same method names among different types.
// + you don't need to type `printGame`, it's just: `print`.
// 타입에 함수를 정의하는 방법
//      리시버
// func (b book) printBook() {
// 	// b is a copy of the original `book` value here.
// 	fmt.Printf("%-15s: $%.2f\n", b.title, b.price)
// }
// 스코프가 타입으로 정의되기 때문에 여러 개의 같은 함수를 한 코드에 저장할 수 있다.
//
// mobydick := book{
// 	title: "moby dick",
// 	price: 10,
// }
// mobydick.print()  으로 사용, 사용 시 mobydick의 복사본이 자동으로 로컬 변수에 할당됨

// ----------------------------------------------------------------------------
// b is a copy of the original `book` value here.
//
// func printBook(b book) {
// 	fmt.Printf("%-15s: $%.2f\n", b.title, b.price)
// }
