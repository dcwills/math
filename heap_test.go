// Copyright (c) 2019 Dean Connable Wills
// All Rights Reserved, except as explicitly granted via the project license.
//
package math

import (
	"testing"
)

// verifies that heapPermutation tries every possible permutation exactly once.
func Test_HeapPermutation(t *testing.T) {
	type TestData struct {
		t *testing.T
		set []bool
	}
	run := func(td interface{}, a []interface{}) {
		testData := td.(*TestData)
		arr := make([]int, len(a)+1)
		for i, v := range a {
			arr[i+1] = v.(int)
		}
		rank := rankParksWills0(arr)
		set := testData.set
		if rank > len(set) || set[rank] {
			testData.t.Errorf("duplicate permutation = %v", rank)
		}
		set[rank] = true
	}

	type args struct {
		run  HeapCallback
		a    []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"random 3",
			args{
				run,
				[]int{3, 1, 2},
			},
		},
		{
			"basic 6",
			args{
				run,
				[]int{1, 2, 3, 4, 5, 6},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := len(tt.args.a)
			array := make([]interface{}, n)
			for i, v := range tt.args.a {
				array[i] = v
			}
			var factorial = []int{1, 1, 2, 6, 24, 120, 720, 5040}
			Heap(&TestData{t, make([]bool, factorial[n])}, tt.args.run, array, n)
		})
	}
}
