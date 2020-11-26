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
		counter int
		V       int  //
		P       *int //int의 포인터를 저장하는 타입
	)

	counter = 100 // counter is an int variable
	P = &counter  // P is a pointer int variable
	V = *P        // P는 주소값, *을 붙여 주소값을 따라간 실제값(counter 출력) V is a int variable (a copy of counter)

	if P == nil {
		fmt.Printf("P is nil and its address is %p\n", P)
		// 0x0 hex 값
	}

	if P == &counter {
		fmt.Printf("P is equal to counter's address: %p == %p\n",
			P, &counter)
		// P is equal to counter's address: 0xc000014090 == 0xc000014090
	}

	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	fmt.Printf("P      : %-13p addr: %-13p *P: %-13d\n", P, &P, *P)
	fmt.Printf("V      : %-13d addr: %-13p\n", V, &V)
	/*
	   counter: 100           addr: 0xc000014090
	   P      : 0xc000014090(주소값)  addr: 0xc000006028(주소값의주소)  *P: 100
	   V      : 100(복사된값)           addr: 0xc000014098(새로운 V변수의 주소)
	*/

	fmt.Println("\n••••• change counter")
	counter = 10 // V doesn't change because it's a copy
	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	fmt.Printf("V      : %-13d addr: %-13p\n", V, &V)
	/*
		••••• change counter
		counter: 10            addr: 0xc000014090
		V      : 100           addr: 0xc000014098
	*/

	fmt.Println("\n••••• change counter in passVal()")
	counter = passVal(counter)
	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	/*
	   ••••• change counter in passVal()
	   함수에 매개변수를 통해 전달한 값은 새로운 주소에 복사된다.
	   리턴값을 counter에 할당하여 함수 내의 n값 50이 대입되었다.
	   n      : 50            addr: 0xc0000140b0
	   counter: 50            addr: 0xc000014090
	*/

	fmt.Println("\n••••• change counter in passPtrVal()")
	passPtrVal(&counter) // same as passPtrVal(&counter) (no need to return)
	passPtrVal(&counter) // same as passPtrVal(&counter) (no need to return)
	fmt.Printf("counter: %-13d addr: %-13p\n", counter, &counter)
	/*
		••••• change counter in passPtrVal()
		포인터를 매개변수로 전달하였다.
		pn     : 0xc000014090  addr: 0xc000006038  *pn: 50
		pn     : 0xc000014090  addr: 0xc000006038  *pn: 51
		pn     : 0xc000014090  addr: 0xc000006040  *pn: 51
		pn     : 0xc000014090  addr: 0xc000006040  *pn: 52
		counter: 52            addr: 0xc000014090
	*/
}

// *pn is a int pointer variable (copy of P)
func passPtrVal(pn *int) {
	fmt.Printf("pn     : %-13p addr: %-13p *pn: %d\n", pn, &pn, *pn)

	// pointers can breach function isolation borders
	*pn++ // counter changes because `pn` points to `counter` — (*pn)++
	fmt.Printf("pn     : %-13p addr: %-13p *pn: %d\n", pn, &pn, *pn)

	// setting it to nil doesn't effect the caller (the main function)
	pn = nil
}

// n is a int variable (copy of counter)
func passVal(n int) int {
	n = 50 // counter doesn't change because `n` is a copy
	fmt.Printf("n      : %-13d addr: %-13p\n", n, &n)
	return n
}
