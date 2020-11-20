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

func main() {
	var (
		sum int
		i   = 1
	)

	//true로 고정되어 있기 때문에 무한반복
	for {
		if i > 5 {
			break
		}

		sum += i

		fmt.Println(i, "-->", sum)

		i++
	}

	fmt.Println(sum)
}
