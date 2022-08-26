package main

import "fmt"

func main() {
	var s *sk
	s = &sk{
		v: 10,
	}
	test(s)
	fmt.Println(s.v)
}

type sk struct {
	v int
}

func test(s *sk) {
	s.v = 5
	fmt.Println(s.v)
}
