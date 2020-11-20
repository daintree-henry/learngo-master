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

	for {
		if i > 10 {
			break
		}

		if i%2 != 0 {
			// just by putting this here we solve the problem
			//이후 코드를 실행하지 않고 바로 다음 분기를 시작한다.
			i++
			continue
		}

		sum += i

		fmt.Println(i, "-->", sum)

		i++
	}

	fmt.Println(sum)
}
