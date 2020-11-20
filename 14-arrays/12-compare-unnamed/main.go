// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// STORY:
// You want to compare two bookcases,
// whether they're equal or not.

func main() {
	type (
		// integer int

		bookcase [5]int
		cabinet  [5]int
		//cabinet  [6]int --> 언더 타입이 달라지기 때문에 타입 전환 불가
		//cabinet  [5]integer --> 언더 타입이 다르다(bookcase는 [5]int , [5]integer->int)

		//          ^- try changing this to: integer
		//             but first: uncomment the integer type above
	)

	blue := bookcase{6, 9, 3, 2, 1}
	red := cabinet{6, 9, 3, 2, 1}

	fmt.Print("Are they equal? ")

	if cabinet(blue) == red {
		fmt.Println("✅")
	} else {
		fmt.Println("❌")
	}

	fmt.Printf("blue: %#v\n", blue)
	fmt.Printf("red : %#v\n", bookcase(red))

	// ------------------------------------------------
	// The underlying type of an unnamed type is itself.
	//
	//   메모리에 실제로 적용되는 타입은 자기 자신이다
	//   [5]integer's underlying type: [5]integer
	//   [5]int's underlying type    : [5]int
	//
	//   > [5]integer and [5]int are different types.
	//   > Their memory layout is not important.
	//   > Their types are not the same.

	// _ = [5]integer{} == [5]int{}

	// ------------------------------------------------
	// An unnamed and a named type can be compared,
	// if they've identical underlying types.
	//
	//   [5]integer's underlying type: [5]integer
	//   cabinet's underlying type   : [5]integer
	//
	// Note: Assuming the cabinet's type definition is like so:
	//       type cabinet [5]integer

	// _ = [5]integer{} == cabinet{}

	// 두 개의 네임드 타입끼리는 언더 타입이 같아도 비교할 수 없다.
	//  bookcase{} == cabinet{} (x)
	// 둘 중의 하나를 변환하여 사용해야 한다
	// bookcase{} == bookcase(cabinet{})
}
