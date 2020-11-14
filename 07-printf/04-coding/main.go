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
	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	// swiss army knife %v verb
	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v millions kms\n", distance)
	fmt.Printf("Orbital Period: %v days\n", orbital)
	fmt.Printf("Does %v have life? %v\n", planet, hasLife)

	// %v 모든 자료형의 값 출력

	// argument indexing - indexes start from 1
	fmt.Printf(
		"%v is %v away. Think! %[2]v kms! %[1]v OMG.\n",
		planet, distance,
	)
	//일반적으로 동사와 변수의 갯수는 일치해야 한다
	//%[1]v, %[2]v 등의 형태는 여러개 사용 가능

	// 타입의 안정성을 위해 v보다 세부적인 동사를 사용하는 것이 좋다.
	// correct verbs:
	// fmt.Printf("Planet: %s\n", planet)
	// fmt.Printf("Distance: %d millions kms\n", distance)
	// fmt.Printf("Orbital Period: %f days\n", orbital)
	// fmt.Printf("Does %s has life? %t\n", planet, hasLife)

	// precision
	// 반올림하는 자리의 위치 결정
	fmt.Printf("Orbital Period: %f days\n", orbital)   //224.701000
	fmt.Printf("Orbital Period: %.0f days\n", orbital) //225
	fmt.Printf("Orbital Period: %.1f days\n", orbital) //224.7
	fmt.Printf("Orbital Period: %.2f days\n", orbital) //224.70
}
