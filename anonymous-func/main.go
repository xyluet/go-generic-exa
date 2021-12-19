package main

import "fmt"

func main() {
	equalFn := makeComparableFunc[int]()
	fmt.Println(equalFn(1, 2))
}

func makeComparableFunc[T comparable]() func(a, b T) bool {
	return func(a, b T) bool {
		return a == b
	}
}
