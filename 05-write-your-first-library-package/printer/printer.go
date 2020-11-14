// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package printer

import "fmt"

// Hello is an exported function
func hello() {
	fmt.Println("exported hello")
}

//go run printer.go
// cannot run non-main package

//라이브러리 패키지는 컴파일할 필요 없이 바로 임포트해서 사용 가능하다
