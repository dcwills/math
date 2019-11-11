// Copyright (c) 2019 Dean Connable Wills
// All Rights Reserved, except as explicitly granted via the project license.
//
package math

type HeapCallback func(interface{}, []interface{})

// Generates permutations using Heap's Algorithm.
// https://en.wikipedia.org/wiki/Heap%27s_algorithm
// This a standard permutation generation algorithm. Array size n generates n! permutations
// we implement it in a way where it can be used to call a function that satisfies permutationTest
//
// this is a *generic* routine and may be reused for many tests.
//
func Heap(data interface{}, run HeapCallback, a []interface{}, size int) {
	// when the size becomes 1, we run the permutationTest with the permuted array
	if size == 1 {
		run(data, a)
		return
	}
	for i := 0; i < size; i++ {
		Heap(data, run, a, size-1)
		if size%2 == 1 {
			a[0], a[size-1] = a[size-1], a[0]
		} else {
			a[i], a[size-1] = a[size-1], a[i]
		}
	}
}
