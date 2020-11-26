// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

func main() {
	mobydick := book{
		title: "moby dick",
		price: 10,
	}

	minecraft := game{
		title: "minecraft",
		price: 20,
	}

	tetris := game{
		title: "tetris",
		price: 5,
	}

	// #3
	mobydick.print()  // sends `mobydick` value to `book.print`
	minecraft.print() // sends `minecraft` value to `game.print`
	tetris.print()    // sends `tetris` value to `game.print`

	// #2 타입에 함수 정의하여 사용
	// mobydick.printBook()
	// minecraft.printGame()

	// #1 일반적인 함수 선언 및 사용
	// printBook(mobydick)
	// printGame(minecraft)
}
