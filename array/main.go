package main

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

type item struct {
	name  string
	price float64
}

func main() {
	fmt.Println(slices.Index([]int{1, 2}, 1))

	// casual
	fmt.Println(Every([]int{2, 4, 6}, isEven))                                                            // true
	fmt.Println(Some([]int{1, 2, 3}, isEven))                                                             // true
	fmt.Println(Some([]int{1, 3, 5}, isEven))                                                             // false
	fmt.Println(Filter([]int{1, 2, 3, 4}, isEven))                                                        // [2, 4]
	fmt.Println(FindIndex([]int{1, 2, 3}, isEven))                                                        // 1
	fmt.Println(FindIndex([]int{1}, isEven))                                                              // -1
	fmt.Println(IndexOf([]int{1, 2}, 2))                                                                  // 1
	fmt.Println(IndexOf([]int{1, 2}, 3))                                                                  // -1
	fmt.Println(Includes([]int{1, 2, 3}, 2))                                                              // true
	fmt.Println(Includes([]int{1, 2, 3}, 4))                                                              // false
	fmt.Println(Join([]int{1, 2}, "-"))                                                                   // 1-2
	fmt.Println(Join([]string{"foo", "bar"}, "-"))                                                        // foo-bar
	fmt.Println(Map([]int{1, 2}, func(i int) int { return i + 10 }))                                      // [11, 12]
	fmt.Println(Map([]string{"foo", "bar"}, func(s string) string { return s + "fazz" }))                 // [foofazz, barfazz]
	fmt.Println(Reduce([]int{1, 2, 3}, func(acc, v int) int { return acc + v }, 0))                       // 6
	fmt.Println(Reduce([]string{"pu", "tang", "ina"}, func(acc, v string) string { return acc + v }, "")) // putangina
	fmt.Println(Reduce([]item{}, func(total float64, current item) float64 { return total + current.price }, 0))

	isEveryEven := MakeEveryFunc(isEven)
	fmt.Println(isEveryEven([]int{2, 4}))
}

func Every[T any](s []T, f func(T) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

func MakeEveryFunc[T any](f func(T) bool) func([]T) bool {
	return func(s []T) bool { return Every(s, f) }
}

func isEven(i int) bool { return i%2 == 0 }

func Some[T any](s []T, f func(T) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

func Filter[T any, S ~[]T](s S, f func(T) bool) S {
	ns := make([]T, 0)
	for _, v := range s {
		if f(v) {
			ns = append(ns, v)
		}
	}
	return ns
}

func FindIndex[T any](s []T, f func(T) bool) int {
	for i, v := range s {
		if f(v) {
			return i
		}
	}
	return -1
}

func IndexOf[T comparable](s []T, e T) int {
	return FindIndex(s, func(e2 T) bool { return e == e2 })
}

func Includes[T comparable](s []T, e T) bool {
	return IndexOf(s, e) >= 0
}

func Join[T any](s []T, sep string) string {
	ss := make([]string, 0, len(s))
	for _, v := range s {
		ss = append(ss, fmt.Sprint(v))
	}
	return strings.Join(ss, sep)
}

func Map[T any, S ~[]T](s S, f func(T) T) S {
	s2 := make(S, 0)
	for _, v := range s {
		s2 = append(s2, f(v))
	}
	return s2
}

type Operation[T, R any] func(R, T) R

func Reduce[T, R any](s []T, f Operation[T, R], initial R) R {
	accumulator := initial
	for _, v := range s {
		accumulator = f(accumulator, v)
	}
	return accumulator
}
