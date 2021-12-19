package main

import "fmt"

type myInt int

type Comparable interface {
	~int | float64 | rune
}

func main() {
	fmt.Println(Equal(1, 1))
	fmt.Println(Equal(myInt(1), myInt(1)))
	fmt.Println(Equal('a', 'a'))
}

func Equal[T Comparable](a, b T) bool {
	return a == b
}
