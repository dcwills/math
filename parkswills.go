// Copyright (c) 2019 Dean Connable Wills
// All Rights Reserved, except as explicitly granted via the project license.
//
package math

// implements the Parks-Wills ranking of a permutation.
// this maps a permutation of 0, ..., n-1 to x \in Z, where 0 <= x < n!
func rankParksWills(arr []int) int {
	r := 0
	for k := len(arr) - 1; k > 0; k-- {
		i := arr[k]
		for i > k {
			i = arr[i]
		}
		arr[k] = i // makes this O(n) and not O(n^2)
		r = (k - i) + (k+1)*r
	}
	return r
}
