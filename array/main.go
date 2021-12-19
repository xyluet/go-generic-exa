package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	fmt.Println(slices.Index([]int{1, 2}, 1))

	// casual
	fmt.Println(Every([]int{2, 4, 6}, isEven))
	fmt.Println(Filter([]int{1, 2, 3, 4}, isEven))
	fmt.Println(FindIndex([]int{1, 2, 3}, isEven))
	fmt.Println(FindIndex([]int{1}, isEven))
	fmt.Println(IndexOf([]int{1, 2}, 2))
	fmt.Println(IndexOf([]int{1, 2}, 3))

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

func Filter[E any, S ~[]E](s []E, f func(e E) bool) S {
	ns := make([]E, 0)
	for _, v := range s {
		if f(v) {
			ns = append(ns, v)
		}
	}
	return ns
}

func FindIndex[E comparable](s []E, f func(e E) bool) int {
	for i, v := range s {
		if f(v) {
			return i
		}
	}
	return -1
}

func IndexOf[E comparable](s []E, e E) int {
	return FindIndex(s, func(e2 E) bool { return e == e2 })
}
