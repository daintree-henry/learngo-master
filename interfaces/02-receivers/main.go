// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

func main() {
	var (
		mobydick  = book{title: "moby dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)

	// sends a pointer of minecraft to the discount method
	// same as: (&minecraft).discount(.1)
	minecraft.discount(.1) // discount가 포인터를 받기 때문에 go가 자동으로 처리한다

	mobydick.print()
	//game.print(minecraft)
	minecraft.print()

	// game.print(tetris)
	// func print(b book) {
	// 	fmt.Printf("%-15s: $%.2f\n", b.title, b.price)
	// } 와 같다
	// == 메소드는 리시버를 첫 번째 파라미터로 받는 함수이다.
	// 메소드는 타입에 속하고, 함수는 패키지에 속한다.
	tetris.print()

	// creates a variable that occupies 200mb on memory
	var h huge
	for i := 0; i < 10; i++ {
		h.addr()
	}
}
