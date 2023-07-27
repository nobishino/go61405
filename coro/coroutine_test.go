package coro_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/nobishino/go61405/coro"
)

func Example() {
	// In = int, Out = string
	resume := coro.New(func(in int, yield func(string) int) string{
		for range 2 {
		out := strconv.Itoa(2 * in)
		in = yield(out)
		}
		return "The coroutine has finished."
	})
	fmt.Printf("resume type = %T\n", resume)
	fmt.Println(resume(1))
	fmt.Println(resume(2))
	fmt.Println(resume(3))
	// fmt.Println(resume(4)) This will cause deadlock!
	// Output:
	// resume type = func(int) string
	// 2
	// 4
	// The coroutine has finished.
}

func TestCheckEquality(t *testing.T) {
	testcases := []struct{
		name string
		left []int
		right []int
		expect bool
	}{
		{
			name: "equal",
			left: []int{1,3,2},
			right: []int{1,3,2},
			expect: true,
		},
		{
			name: "not_equal",
			left: []int{1,3,2,1},
			right: []int{1,3,2,4},
			expect: false,
		},
	}

	for _,tc:=range testcases {
		t.Run(tc.name, func(t *testing.T){
			leftIt := iterator(t, tc.left)
			rightIt := iterator(t, tc.right)

			type token = struct{} // token = In, int = out
			tk := token{}

			leftResume := coro.New(func(in token, yield func(int)token) int{
				for v := range leftIt {
					yield(v)
				}
				return 0
			})

			rightResume := coro.New(func(in token, yield func(int)token) int{
				for v := range rightIt {
					yield(v)
				}
				return 0
			})

			result := true
			for {
				l := leftResume(tk)
				r := rightResume(tk)
				if l != r {
					result = false
					break
				}
				if l == 0 {
					break
				}
			}

			if result != tc.expect {
				t.Errorf("expect %t but got %t", tc.expect, result)
			}
		})
	}
}

// 値のみのスライスから、それをyieldするiteratorを返す
func iterator[V any](t *testing.T, xs []V) func(yield func(V) bool) bool {
	t.Helper()
	return func(yield func(V) bool) bool {
		for _, x := range xs {
			if !yield(x) {
				return false
			}
		}
		return true
	}
}