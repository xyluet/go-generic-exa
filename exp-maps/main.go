package main

import (
	"fmt"

	"golang.org/x/exp/maps"
)

func main() {
	m := map[int]interface{}{
		123: "foo",
		456: "bar",
	}

	//
	fmt.Println(Keys(m))
	fmt.Println(Values(m))
	mm2 := Clone(m)
	fmt.Println(mm2)
	fmt.Println(Equal(m, mm2))

	//
	fmt.Println(maps.Keys(m))
	fmt.Println(maps.Values(m))
	m2 := maps.Clone(m)
	fmt.Println(m2)
	fmt.Println(maps.Equal(m, m2))
}

func Keys[M map[K]V, K comparable, V any](m M) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func Values[M map[K]V, K comparable, V any](m M) []V {
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func Clone[M map[K]V, K comparable, V any](m M) M {
	r := make(M, len(m))
	for k, v := range m {
		r[k] = v
	}
	return r
}

func Equal[K, V comparable, M map[K]V](m1, m2 M) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}
