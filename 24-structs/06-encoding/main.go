// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool // #3

type user struct { // #1
	//필드 태그: 타입의 메타데이터, 키밸류 형태의 값이며 인코딩,디코딩에 주로 사용됨

	Name        string      `json:"username"`        //표현되는 인코딩 이름을 바꿀 수 있다
	Password    string      `json:"-"`               //필드태그를 통해 특정 타입을 제거할 수 있다.
	Permissions permissions `json:"perms,omitempty"` //optteempty: nil인 자료를 제외한다.

	// name        string // #1
	// password    string // #1
	// permissions // #3
	// 위의 형태로 사용하면 아무것도 출력되지 앟는다..
	// json 패키지는 exported 필드만 사용 가능하다(대문자)
}

func main() {
	users := []user{ // #2
		{"inanc", "1234", nil},
		{"god", "42", permissions{"admin": true}},
		{"devil", "66", permissions{"write": true}},
	}

	// out, err := json.Marshal(users) // #4
	out, err := json.MarshalIndent(users, "", "\t") // #5
	if err != nil {                                 // #4
		fmt.Println(err)
		return
	}

	fmt.Println(string(out)) // #4
}
