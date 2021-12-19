package main

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

func main() {
	fmt.Println(slices.Index([]int{1, 2}, 1))

	// casual
	fmt.Println(Every([]int{2, 4, 6}, isEven))                                            // true
	fmt.Println(Filter([]int{1, 2, 3, 4}, isEven))                                        // [2, 4]
	fmt.Println(FindIndex([]int{1, 2, 3}, isEven))                                        // 1
	fmt.Println(FindIndex([]int{1}, isEven))                                              // -1
	fmt.Println(IndexOf([]int{1, 2}, 2))                                                  // 1
	fmt.Println(IndexOf([]int{1, 2}, 3))                                                  // -1
	fmt.Println(Includes([]int{1, 2, 3}, 2))                                              // true
	fmt.Println(Includes([]int{1, 2, 3}, 4))                                              // false
	fmt.Println(Join([]int{1, 2}, "-"))                                                   // 1-2
	fmt.Println(Join([]string{"foo", "bar"}, "-"))                                        // foo-bar
	fmt.Println(Map([]int{1, 2}, func(i int) int { return i + 10 }))                      // [11, 12]
	fmt.Println(Map([]string{"foo", "bar"}, func(s string) string { return s + "fazz" })) // [foofazz, barfazz]

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

func Includes[E comparable](s []E, e E) bool {
	return IndexOf(s, e) >= 0
}

func Join[E any](s []E, sep string) string {
	ss := make([]string, 0, len(s))
	for _, v := range s {
		ss = append(ss, fmt.Sprint(v))
	}
	return strings.Join(ss, sep)
}

func Map[E any, S ~[]E](s S, f func(e E) E) S {
	s2 := make(S, 0)
	for _, v := range s {
		s2 = append(s2, f(v))
	}
	return s2
}
