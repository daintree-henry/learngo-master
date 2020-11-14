// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

/*
모든 패키지 파일은 하나의 디렉토리에 존재해야 한다.
모든 파일은 같은 패키지를 가지고 있어야 한다
 = 같은 패키지에 있는 파일은 하나의 파일로 생각할 수 있다.

하나의 파일은 하나의 패키지만 가질 수 있다.

패키지는 실행 패키지와 라이브러리 패키지가 있다.

실행 패키지 : 실행을 위한 패키지
- main 패키지에 속해 실행 가능한 go 프로그램
- 임포트 불가능
- main만 사용가능

라이브러리 패키지 : 재사용을 위한 패키지
- 실행 불가
- 임포트하여 사용
- 어떤 이름이던지 사용 가능
- main 함수를 가질 수 없음

*/

package main

import "fmt"

func main() {
	fmt.Println("Hello!")

	// You can access functions from other files
	// which are in the same package

	// Here, `main()` can access `bye()` and `hey()`

	// It's because bye.go, hey.go and main.go
	//   are in the main package.

	bye()
	hey()
}
