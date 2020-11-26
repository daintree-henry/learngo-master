// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

func showN() {
	if N == 0 {
		return
	}
	fmt.Printf("showN       : N is %d\n", N)
}

func incrN() {
	N++
}

// 함수는 0개이상의 매개변수를 가질 수 있다.
// return : 함수 종료
// 여러개의 변수를 리턴할 수 있다
// func AtoI(s string) (int, error){
//   매개변수 s 는 로컬스코프의 변수에 복사된다.
//   return 0,nil <- 반드시 갯수가 맞아야 한다.
//}
//

// you cannot declare a function within the same package with the same name
// func incrN() {
// }

// 이 예제의 경우 패키지레벨 변수인 N이 어떤 함수에 의해 변화하는지 관찰할 수 없기 때문에
// 좋은 사례가 아니다.
