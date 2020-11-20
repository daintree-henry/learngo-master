// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	s "github.com/inancgumus/prettyslice"
)

// Array, Slice, String
// newslice := sliceable[시작인덱스:끝인덱스+1]
func main() {
	msg := []byte{'h', 'e', 'l', 'l', 'o'}
	s.Show("msg", msg)

	s.Show("msg[0:1]", msg[0:1]) // h
	s.Show("msg[0:2]", msg[0:2]) // he
	s.Show("msg[0:3]", msg[0:3]) // hel
	s.Show("msg[0:4]", msg[0:4]) // hell
	s.Show("msg[0:5]", msg[0:5]) // hello

	// default indexes
	s.Show("msg[0:]", msg[0:]) // hello
	s.Show("msg[:5]", msg[:5]) // hello
	s.Show("msg[:]", msg[:])   //

	// error: beyond
	// s.Show("msg", msg)[:6]

	s.Show("msg[1:4]", msg[1:4])

	s.Show("msg[1:5]", msg[1:5])
	s.Show("msg[1:]", msg[1:])

	s.Show("append(msg)", append(msg[:4], '!'))
}
