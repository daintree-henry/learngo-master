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
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}
	query := args[0]

	dict := map[string]string{
		"good":    "iyi",
		"great":   "harika",
		"perfect": "mükemmel",
		"awesome": "mükemmel", // #5
	}

	// turkish := dict // #1
	// turkish["good"] = "güzel"
	// dict["great"] = "kusursuz"
	// 포인터를 사용하기 때문에 모두 변경
	// 같은 맵 헤더의 포인터 주소
	// 맵 헤더는 각 키&밸류의 포인터 주소를 가진 주소록

	delete(dict, "awesome")     // #6
	delete(dict, "awesome")     // #7: no-op
	delete(dict, "notexisting") // 없으면 아무 일도 일어나지 않는다.

	// dict = nil // #8
	for k := range dict { // #9
		delete(dict, k)
	}

	//새로운 맵 생성
	turkish := make(map[string]string, len(dict)) // #3
	for k, v := range dict {
		turkish[v] = k
	}

	// fmt.Printf("english: %q\nturkish: %q\n", dict, turkish) // #4
	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	if value, ok := turkish[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	fmt.Printf("%q not found.\n", query)
}
