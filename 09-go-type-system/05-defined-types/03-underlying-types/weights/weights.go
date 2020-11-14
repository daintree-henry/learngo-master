// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package weights

type (
	//대문자로 시작할 경우 외부에서 참고 가능
	//소문자로 시작할 경우 불가능
	// Gram underlying type is int64
	Gram int

	// Kilogram underlying type is int64
	Kilogram Gram

	// Ton underlying type is int64
	Ton Kilogram
)
