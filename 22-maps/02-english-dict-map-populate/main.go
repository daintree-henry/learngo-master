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
	// #2A: Get the key from CLI
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[english word] -> [turkish word]")
		return
	}
	query := args[0]

	// #1: Empty Map Literal
	// dict := map[string]string{}

	// 3: Map Literal
	dict := map[string]string{
		"good":    "kötü",
		"great":   "harika",
		"perfect": "mükemmel",
		// #4
		// 42: "forty two",
		// "forty two": 42,
	}

	dict["up"] = "yukarı"  // adds a new pair
	dict["down"] = "aşağı" // adds a new pair
	dict["good"] = "iyi"   // #5: overwrites the value at the key: "good"
	dict["mistake"] = ""   // #8: a key with a zero-value

	// #10: comma ok in a short if
	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}
	fmt.Printf("%q not found.\n", query)

	// fmt.Printf("Zero Value: %#v\n", dict)
	fmt.Printf("# of Keys : %d\n", len(dict))

	// #13: compare a map using its printed output
	// copied := map[string]string{"up": "yukarı", "down": "aşağı",
	// 	"mistake": "", "good": "iyi", "great": "harika",
	// 	"perfect": "mükemmel"}
	// first := fmt.Sprintf("%s", dict)
	// second := fmt.Sprintf("%s", copied)

	// if first == second {
	// 	fmt.Println("Maps are equal")
	// }

	// #12: printing a map (ordered output since Go 1.12)
	// 맵 프린트
	// fmt.Printf("%#v\n", dict)

	// #11
	// 맵 순환
	// 출력 시마다 순서가 변환된다(맵은 검색에 특화, 루핑에는 부적절)
	// for k, v := range dict {
	// 	fmt.Printf("%q means %#v\n", k, v)
	// }

	// #9: check for non-existing key: with comma, ok
	// value, ok := dict[query]
	// 선언 시 두 변수를 활용할 경우 두 번째 값은 bool로 키 존재 여부를 반환한다.
	// if !ok {
	// 	fmt.Printf("%q not found.\n", query)
	// }

	// #7: check for non-existing key using zero-value
	// if value == "" {
	// 	fmt.Printf("%q not found.\n", query)
	//  없는 키를 입력하면 nil 값을 출력한다.
	// }

	// #6: getting values from a map using keys directly
	// 맵 값 출력
	// fmt.Println("good      -> ", dict["good"])
	// fmt.Println("great     -> ", dict["great"])
	// fmt.Println("perfect   -> ", dict["perfect"])

	// #2B: retrieve values by key - O(1) efficiency
	// value := dict[query]

}
