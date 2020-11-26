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
	"strings"
)

func main() {
	fmt.Println("••••• ARRAYS")
	arrays()

	fmt.Println("\n••••• SLICES")
	slices()

	fmt.Println("\n••••• MAPS")
	maps()

	fmt.Println("\n••••• STRUCTS")
	structs()
}

// ••••••••••••••••••••••••••••••••••••••••••••••••••

type house struct {
	name  string
	rooms int
}

func structs() {
	myHouse := house{name: "My House", rooms: 5}

	addRoom(myHouse)

	// fmt.Printf("%+v\n", myHouse)
	fmt.Printf("structs()     : %p %+v\n", &myHouse, myHouse)

	addRoomPtr(&myHouse)
	fmt.Printf("structs()     : %p %+v\n", &myHouse, myHouse)
	/*
			••••• STRUCTS
		addRoom()     : 0xc000098540 {name:My House rooms:6}
		structs()     : 0xc000098520(복사) {name:My House rooms:5}
		값이 복사되었기 때문에 변경사항이 반영되지 않음

		addRoomPtr()  : 0xc000098520 &{name:My House rooms:6}
		&h.name       : 0xc000098520
		&h.rooms      : 0xc000098530
		structs()     : 0xc000098520 {name:My House rooms:6}
	*/
}

func addRoomPtr(h *house) {
	h.rooms++ // same: (*h).rooms++
	fmt.Printf("addRoomPtr()  : %p %+v\n", h, h)
	fmt.Printf("&h.name       : %p\n", &h.name)
	fmt.Printf("&h.rooms      : %p\n", &h.rooms)
}

func addRoom(h house) {
	h.rooms++
	fmt.Printf("addRoom()     : %p %+v\n", &h, h)
}

// ••••••••••••••••••••••••••••••••••••••••••••••••••

func maps() {
	confused := map[string]int{"one": 2, "two": 1}
	fix(confused)
	fmt.Println(confused)

	// &confused["one"] 맵 값의 주소는 변경불가능하다. 임의로 조작 불가능하도록 막혀있다.
	/*
		//맵의 값이 포인터이기 때문에 변경이 적용된다.
		••••• MAPS
		map[one:1 three:3 two:2]
	*/
}

func fix(m map[string]int) {
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
}

// ••••••••••••••••••••••••••••••••••••••••••••••••••

func slices() {
	dirs := []string{"up", "down", "left", "right"}

	up(dirs)
	fmt.Printf("slices list   : %p %q\n", &dirs, dirs)

	upPtr(&dirs)
	fmt.Printf("slices list   : %p %q\n", &dirs, dirs)
	/*
			••••• SLICES
		up.list[0]    : 0xc0000d6000
		up.list[1]    : 0xc0000d6010
		up.list[2]    : 0xc0000d6020
		up.list[3]    : 0xc0000d6030
		up list       : 0xc000098460 ["UP" "DOWN" "LEFT" "RIGHT" "HEISEN BUG"]
		slices list   : 0xc000098440 ["UP" "DOWN" "LEFT" "RIGHT"]
		// 헤더가 전달된다.
		// 변경사항이 반영되지만 추가사항은 반영되지 않는다. 슬라이스의 값이 전달되지만 값이 포인터이기 때문이다.
		// 변경사항은 값에 대한 변경사항이 반영되고 len, cap은 변경되지 않는다.
		upPtr list    : 0xc000098440 &["UP" "DOWN" "LEFT" "RIGHT" "HEISEN BUG"]
		slices list   : 0xc000098440 ["UP" "DOWN" "LEFT" "RIGHT" "HEISEN BUG"]
		// 헤더의 포인터가 전달된다.
	*/
}

func upPtr(list *[]string) {
	lv := *list

	for i := range lv {
		lv[i] = strings.ToUpper(lv[i])
	}

	*list = append(*list, "HEISEN BUG")

	fmt.Printf("upPtr list    : %p %q\n", list, list)
}

func up(list []string) {
	for i := range list {
		list[i] = strings.ToUpper(list[i])
		fmt.Printf("up.list[%d]    : %p\n", i, &list[i])
	}

	list = append(list, "HEISEN BUG")

	fmt.Printf("up list       : %p %q\n", &list, list)
}

// ••••••••••••••••••••••••••••••••••••••••••••••••••

func arrays() {
	nums := [...]int{1, 2, 3}

	incr(nums)
	fmt.Printf("arrays nums   : %p\n", &nums)
	fmt.Println(nums)

	incrByPtr(&nums)
	fmt.Println(nums)
	/*
			••••• ARRAYS
		incr nums     : 0xc0000a01a0
		incr.nums[0]  : 0xc0000a01a0
		incr.nums[1]  : 0xc0000a01a8
		incr.nums[2]  : 0xc0000a01b0
		arrays nums   : 0xc0000a0180
		[1 2 3]
		incrByPtr nums: 0xc0000cc020
		[2 3 4]
	*/
}

func incr(nums [3]int) {
	fmt.Printf("incr nums     : %p\n", &nums)
	for i := range nums {
		nums[i]++
		fmt.Printf("incr.nums[%d]  : %p\n", i, &nums[i])
	}
}

func incrByPtr(nums *[3]int) {
	fmt.Printf("incrByPtr nums: %p\n", &nums)
	for i := range nums {
		nums[i]++ // same: (*nums)[i]++
	}
}
