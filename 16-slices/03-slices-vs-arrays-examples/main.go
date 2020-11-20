// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import "fmt"

// STORY:
// You want to list the books and games you have

func main() {
	var books [5]string
	books[0] = "dracula"
	books[1] = "1984"
	books[2] = "island"

	newBooks := [5]string{"ulysses", "fire"}
	if books == newBooks {
	}
	books = newBooks // {"ulysses", "fire", " ", " ", " "} 가 됨 books 배열에 games 슬라이스의 주소값 대입

	games := []string{"kokemon", "sims"}           // 슬라이스 선언 및 초기화
	newGames := []string{"pacman", "doom", "pong"} // 슬라이스 선언 및 초기화
	newGames = games                               // newGames 슬라이스에 games 슬라이스의 주소값 대입, 길이 및 값이 모두 같아짐
	games = nil
	games = []string{}

	var ok string
	for i, game := range games {
		if game != newGames[i] {
			ok = "not "
			break
		}
	}

	if len(games) != len(newGames) {
		ok = "not "
	}

	fmt.Printf("games and newGames are %sequal\n\n", ok)

	fmt.Printf("books         : %#v\n", books)
	fmt.Printf("new books     : %#v\n", newBooks)
	fmt.Printf("games         : %T\n", games)
	fmt.Printf("new games     : %#v\n", newGames)
	fmt.Printf("games's length: %d\n", len(games))
	fmt.Printf("games's nil   : %t\n", games == nil)
}
