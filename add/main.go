package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add("foo", "bar"))
	fmt.Println()
}

func add[T int | string](a, b T) T {
	return a + b
}
