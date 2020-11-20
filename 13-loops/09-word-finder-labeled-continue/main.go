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
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]

queries:
	for _, q := range query {
		for i, w := range words {
			if q == w {
				fmt.Printf("#%-2d: %q\n", i+1, w)

				// 뒤에 레이블이 없을 경우 자신이 속한 for루프로 이동
				// skip duplicate words
				continue queries
			}
		}
	}
}
