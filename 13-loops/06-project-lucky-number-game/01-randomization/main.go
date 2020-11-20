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
	"math/rand"
	"time"
)

func main() {
	//시드넘버에 따라 랜덤 패턴이 달라진다
	// rand.Seed(10)
	// rand.Seed(100)

	// t := time.Now()
	// rand.Seed(t.UnixNano())
	// 시간정보를 시드에 넣어 완전한 랜덤 넘버를 획득한다

	// ^-- same:

	rand.Seed(time.Now().UnixNano())

	guess := 10

	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
