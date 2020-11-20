// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	{
		// its length is part of its type
		// []의 길이는 컴파일 타임에 속하기 때문에 런타임에서 변경이 불가능하다
		// 따라서 동적으로 변경 불가
		var nums [5]int
		// 선언시 모든 값이 자료형의 기본값으로 변경된다, lenth는 5로 고정됨
		// { 0, 0, 0, 0, 0 }
		fmt.Printf("nums array: %#v\n", nums)
	}

	{
		// its length is not part of its type
		// 비어있으면 슬라이스이다.
		// 선언 시에는 값이 없이 비어있게 된다. nil, lenth는 0
		var nums []int

		fmt.Printf("nums slice: %#v\n", nums)

		fmt.Printf("len(nums) : %d\n", len(nums))

		// won't work: the slice is nil.
		// fmt.Printf("nums[0]: %d\n", nums[0])
		// fmt.Printf("nums[1]: %d\n", nums[1])
	}
}
