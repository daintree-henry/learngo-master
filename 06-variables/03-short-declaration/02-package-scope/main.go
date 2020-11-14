// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// You can't use declaration statements without a keyword
// Short declaration doesn't have a keyword (`var`)
// So, it can't be used at the package scope
// 패키지 스코프에는 := 형태로 변수 선언이 불가능하다.

//
// SYNTAX ERROR:
// "non-declaration statement outside function body"

// safe := true

// However, you can use the normal declaration at the
// package scope. Since it has a keyword: `var`

// 패키지 변수는 프로그램이 종료될 때가지 남아있다.
// 패키지 스코프에서는 키워드로 시작해야 한다. (ex package, var, func)
var safe = true

func main() {
	fmt.Println(safe)
}
