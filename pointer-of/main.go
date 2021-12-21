package main

import "fmt"

func main() {
	a := PointerOf(10)
	fmt.Println(a) // address a
	b := a
	*a = 5
	fmt.Println(*a, *b)
}

func PointerOf[T any](v T) *T { return &v }
