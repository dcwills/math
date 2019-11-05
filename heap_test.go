// Copyright (c) 2019 Dean Connable Wills
// All Rights Reserved, except as explicitly granted via the project license.
//
package math
import "testing"

func TestHeap(t *testing.T) {
    type args struct {
        t    interface{}
        run  HeapCallback
        a    []interface{}
        size int
    }
    tests := []struct {
        name string
        args args
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
        })
    }
}

func Test_rankParksWills0(t *testing.T) {
    type args struct {
        arr []int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := rankParksWills0(tt.args.arr); got != tt.want {
                t.Errorf("rankParksWills0() = %v, want %v", got, tt.want)
            }
        })
    }
}