// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

type money float64

func (m money) string() string {
	// $xx.yy
	return fmt.Sprintf("$%.2f", m)
}

//베이직 타입에 함수를 붙여서 활용할 수 있다.
