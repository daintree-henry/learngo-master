// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"sort"

	s "github.com/inancgumus/prettyslice"
)

func main() {
	// #1: arrays and non-empty slice literals create an array.
	// For the arrays, it's explicit, but for the slices,
	// it's done implicitly, behind the scenes.

	grades := [...]float64{40, 10, 20, 50, 70, 60} // #1
	// #2: cheap: slicing doesn't allocate new memory (array).
	front := grades[:3]
	// front와 grades는 같은 backing array를 갖는다.

	// grades := []float64{40, 10, 20, 50, 60, 70} // #4
	// let's break the connection
	// var newGrades []float64
	// newGrades = append(newGrades, grades...)
	// 새로운 슬라이스 생성 후 모든 값을 더한다
	// front := newGrades[:3]

	// #6: shortcut: []float64(nil) is a nil float64 slice
	// newGrades := append([]float64(nil), grades...)

	// #3: sort its first segment
	sort.Float64s(front)
	// backing array의 3번째 부분까지만 정렬한다

	// #7: new slices look at the same backing array
	// front, front2, front3, newGrades, they all have the same backing array
	// front2 := front[:3]
	// front3 := front

	s.PrintBacking = true       // #1
	s.MaxPerLine = 7            // #1
	s.Show("grades", grades[:]) // #1
	// s.Show("newGrades", newGrades) // #5
	s.Show("front", front) // #2
	// s.Show("front2", front2)       // #7
	// s.Show("front3", front3)       // #7
}
