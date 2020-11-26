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
	"unsafe"
)

//String은 rune형의 리스트이다.
//String은 String Header(Pointer, Length)를 갖는다
//Pointer는 읽기전용 backing array의 주소이다.
//read-only이기 때문에 capacity 값이 없다.

func main() {
	// empty := ""
	// dump(empty)

	hello := "hello"
	dump(hello)
	dump("hello")  // 같은 문자열은 같은 backing array를 공유한다.
	dump("hello!") // 다른 backing array를 가진다.

	for i := range hello {
		dump(hello[i : i+1])
	}

	dump(string([]byte(hello)))
	dump(string([]byte(hello)))
	dump(string([]rune(hello))) // 모두 다른 backing array를 갖는다.
}

// StringHeader is used by a string value
// In practice, you should use: reflect.Header
type StringHeader struct {
	// points to a backing array's item
	pointer uintptr // where it starts
	length  int     // where it ends
}

// dump prints the string header of a string value
// 스트링헤더를 출력하는 함수
func dump(s string) {
	ptr := *(*StringHeader)(unsafe.Pointer(&s))
	fmt.Printf("%q: %+v\n", s, ptr)
}
