package myrange_test

import (
	"fmt"

	"github.com/nobishino/go61405/myrange"
)

func ExampleBuiltinRange0() {
	it := myrange.Iterator0(10)
	var i int
	for range it {
		i++
		if i < 2 {
			continue
		}
		fmt.Println(i)
		if i > 3 {
			break
		}
	}
	// Output:
	// 2
	// 3
	// 4
}

func ExampleMyRange0() {
	it := myrange.Iterator0(10)
	var i int
	myrange.MyRange0(it, func()bool {
		i++
		if i < 2 {
			return true
		}
		fmt.Println(i)
		if i > 3 {
			return false
		}
		return true
	})
	// Output:
	// 2
	// 3
	// 4
}

func ExampleBuiltinRange1() {
	it := myrange.Iterator1([]int{
		1,2,4,8,16,32,
	})
	for v:= range it {
		if v < 4 {
			continue
		}
		fmt.Println(v)
		if v > 10 {
			break
		}
	}
	// Output:
	// 4
	// 8
	// 16
}

func ExampleMyRange1() {
	it := myrange.Iterator1([]int{
		1,2,4,8,16,32,
	})
	myrange.MyRange1(it, func(v int)bool{
		if v < 4 {
			return true // continue
		}
		fmt.Println (v)
		if v > 10 {
			return false // break
		}
		return true
	})
	// Output:
	// 4
	// 8
	// 16
}

func ExampleBuiltinRange2() {
	it := myrange.Iterator2([]myrange.Entry[string, int]{
		{"alpha", 5},
		{"beta", 4},
		{"x", 1},
		{"gamma", 5},
	})
	for k, v := range it {
		fmt.Println(k, v)
		if v == 1 {
			break
		}
	}
	// Output:
	// alpha 5
	// beta 4
	// x 1
}

// ExampleBuiltinRangeの書き換え
func ExampleMyRange2() {
	it := myrange.Iterator2([]myrange.Entry[string, int]{
		{"alpha", 5},
		{"beta", 4},
		{"x", 1},
		{"gamma", 5},
	})
	myrange.MyRange2(it, func(key string, val int) bool {
		fmt.Println(key, val) // statement
		if val == 1 {
			return false // break
		}
		return true // continue
	})
	// Output:
	// alpha 5
	// beta 4
	// x 1
}
