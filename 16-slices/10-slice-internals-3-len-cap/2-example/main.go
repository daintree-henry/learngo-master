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

func main() {
	s.PrintBacking = true

	// ----------------------------------------------------
	// #1 nil slice
	var games []string // nil slice
	s.Show("games", games)
	// 3개 요소의 모든 값 0

	// ----------------------------------------------------
	// #2 empty slice
	games = []string{} // empty slice
	s.Show("games", games)
	// s.Show("another empty", []int{}) //empty 슬라이스는 같은 포인터를 참조
	// 초기화 후 값:
	// len:0, cap:0, ptr:포인터값

	// ----------------------------------------------------
	// #3 non-empty slice
	games = []string{"pacman", "mario", "tetris", "doom"}
	s.Show("games", games)
	// len:4, cap:4, ptr: empty상태와 다른 포인터

	// ----------------------------------------------------
	// #4 reset the part using the games slice
	//    part is empty but its cap is still 4
	part := games //새로운 슬라이스 헤더 생성
	s.Show("part", part)
	// games와 같은 포인터

	part = games[:0]
	s.Show("part[:0]", part)
	//len:0, cat:4, ptr: games와 같은 포인터
	s.Show("part[:cap]", part[:cap(part)]) // 리슬라이스
	//len:4, cat:4, ptr: games와 같은 포인터

	for cap(part) != 0 {
		part = part[1:cap(part)]
		s.Show("part", part)
	}
	/*
		part[:cap]          (len:4  cap:4  ptr:6656)
		part                (len:3  cap:3  ptr:6672)
		part                (len:2  cap:2  ptr:6688)
	*/

	// #6 backing array's elements become inaccessible
	// games = games[len(games):]

	// ----------------------------------------------------
	// #5 part doesn't have any more capacity
	//    games slice is still intact
	s.Show("part", part)
	s.Show("games", games)
}
