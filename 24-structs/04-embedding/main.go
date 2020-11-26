// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"fmt"
)

//임베딩 : 타입을 중첩적으로 사용한다.
func main() {
	type text struct {
		title string
		words int
	}

	type book struct {
		// #3: include the text as a field
		// 일반적인 선언
		// text text   -> text.title, text.words로 접근해야함

		// #4: embed the text
		// 익명필드(anonymous field)로 사용하는 경우
		// text 구조체의 title,words 변수를 직접 접근하여 사용할 수 있다.
		// .text.title , .text.words 형르로도 사용가능
		text
		isbn string

		// #5: add a conflicting field
		// 익명필드와 동일한 변수명을 선언할 경우
		// .title로 호출할 경우 임베디드 타입보다 우선권을 갖는다.
		title string
	}

	// #2: print a book
	// moby := book{title: "moby dick", words: 206052, isbn: "102030"}
	// fmt.Printf("%s has %d words (isbn: %s)\n", moby.title, moby.words, moby.isbn)

	// #3b: type the text in its own field
	moby := book{
		// #5c: type the field in a new field
		// title: "conflict",
		text: text{title: "moby dick", words: 206052},
		isbn: "102030",
	}

	moby.text.words = 1000
	moby.words++

	// // #4b: print the book
	fmt.Printf("%s has %d words (isbn: %s)\n",
		moby.title, // equals to: moby.text.title
		moby.words, // equals to: moby.text.words
		moby.isbn)

	// #3c: print the book
	// fmt.Printf("%s has %d words (isbn: %s)\n",
	// 	moby.text.title, moby.text.words, moby.isbn)

	// #5b: print the conflict
	fmt.Printf("%#v\n", moby)

	// go get -u github.com/davecgh/go-spew/spew
	// spew.Dump(moby)
}
