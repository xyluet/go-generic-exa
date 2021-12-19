package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	fmt.Println(slices.Index([]int{1, 2}, 1))

	// casual
	fmt.Println(Every([]int{2, 4, 6}, isEven))

	isEveryEven := MakeEveryFunc(isEven)
	fmt.Println(isEveryEven([]int{2, 4}))
}

func Every[E any](s []E, f func(E) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

func MakeEveryFunc[E any](f func(e E) bool) func(s []E) bool {
	return func(s []E) bool { return Every(s, f) }
}

func isEven(i int) bool { return i%2 == 0 }
