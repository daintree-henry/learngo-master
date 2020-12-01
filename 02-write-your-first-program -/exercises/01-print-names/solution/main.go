// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// 모든 코드는 package로 시작해야 한다
// 파일이 속한 패키지의 이름을 기재한다.
package main

import "fmt"

/*

go build
main.go 파일을 exe파일로 빌드한다.
파일명=상위폴더명
소스코드에서 go build명령어를 통한 컴파일링을 통해 컴파일 코드가 됨

go run main.go
컴파일과 파일 실행을 동시에 진행

go run -x main.go
디버깅 가능


./solution
파일 실행
*/

//main함수는 Go 프로그램 실행 시 호출되는 스페셜 함수이다.
func main() {
	fmt.Println("Nikola")
	fmt.Println("Thomas")
}
