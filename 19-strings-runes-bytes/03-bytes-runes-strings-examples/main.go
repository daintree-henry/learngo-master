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
	"unicode/utf8"
	"unsafe"
)

func main() {
	str := "Yūgen ☯ 💀"

	// can't change a string
	// a string is a read-only byte-slice 읽기 전용이기 때문에 수정할 수 없다.
	// str[0] = 'N'
	// str[1] = 'o'

	bytes := []byte(str)
	// 바이트 값으로 변환하여 변경한다.
	// can change a byte slice
	// bytes[0] = 'N'
	// bytes[1] = 'o'

	str = string(bytes)
	//변경된 값 출력

	fmt.Printf("%s\n", str)                                 // 문자열 출력
	fmt.Printf("\t%d bytes\n", len(str))                    // 15 bytes -> rune이 아닌 bytes를 카운트
	fmt.Printf("\t%d runes\n", utf8.RuneCountInString(str)) // 9 runes 출력
	fmt.Printf("% x\n", bytes)                              // 바이트 슬라이스 값 출력
	fmt.Printf("\t%d bytes\n", len(bytes))                  // 15 bytes 출력
	fmt.Printf("\t%d runes\n", utf8.RuneCount(bytes))       // 9 runes 출력
	/*
	           15 bytes
	           9 runes
	   59 c5 ab 67 65 6e 20 e2 98 af 20 f0 9f 92 80
	           15 bytes
	           9 runes
	*/
	fmt.Println()
	for i, r := range str {
		fmt.Printf("str[%2d] = % -12x = %q\n", i, string(r), r)
	}
	/*
	   rune    // bytes       // rune의 utf encoded 값
	   str[ 0] = 59           = 'Y'
	   str[ 1] = c5 ab        = 'ū'
	   str[ 3] = 67           = 'g'
	   str[ 4] = 65           = 'e'
	   str[ 5] = 6e           = 'n'
	   str[ 6] = 20           = ' '
	   str[ 7] = e2 98 af     = '☯'
	   str[10] = 20           = ' '
	   str[11] = f0 9f 92 80  = '💀'
	*/

	fmt.Println()
	fmt.Printf("1st byte   : %c\n", str[0])           // ok
	fmt.Printf("2nd byte   : %c\n", str[1])           // not ok
	fmt.Printf("2nd rune   : %s\n", str[1:3])         // ok
	fmt.Printf("last rune  : %s\n", str[len(str)-4:]) // ok

	// disadvantage: each one is 4 bytes : rune슬라이스는 하나에 4byte 사용, 비효율적이다
	// byte배열 string을 이용하면 더 적은 메모리 사용 가능
	runes := []rune(str)

	fmt.Println()
	fmt.Printf("%s\n", str)
	fmt.Printf("\t%d bytes\n", int(unsafe.Sizeof(runes[0]))*len(runes))
	fmt.Printf("\t%d runes\n", len(runes))

	fmt.Printf("1st rune   : %c\n", runes[0])
	fmt.Printf("2nd rune   : %c\n", runes[1])
	fmt.Printf("first five : %c\n", runes[:5])

	fmt.Println()

	word := "öykü"
	fmt.Printf("%q in runes: %c\n", word, []rune(word))
	fmt.Printf("%q in bytes: % [1]x\n", word)

	fmt.Printf("%s %s\n", word[:2], []byte{word[0], word[1]}) // ö
	fmt.Printf("%c\n", word[2])                               // y
	fmt.Printf("%c\n", word[3])                               // k
	fmt.Printf("%s %s\n", word[4:], []byte{word[4], word[5]}) // ü
}
