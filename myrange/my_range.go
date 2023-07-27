package myrange

type Entry[K, V any] struct {
	Key K
	Val V
}

func Iterator0[Num ~int](upperbound Num) func(yield func() bool) bool {
	return func(yield func() bool) bool {
		for i:=Num(0);i<upperbound;i++ {
			if !yield() {
				return false
			}
		}
		return true
	}
}

// 値のみのスライスから、それをyieldするiteratorを返す
func Iterator1[V any](xs []V) func(yield func(V) bool) bool {
	return func(yield func(V) bool) bool {
		for _, x := range xs {
			if !yield(x) {
				return false
			}
		}
		return true
	}
}

// key-valueのsliceからそれをyieldするiteratorを返す
func Iterator2[K, V any](xs []Entry[K, V]) func(yield func(K, V) bool) bool {
	return func(yield func(K, V) bool) bool {
		for _, x := range xs {
			if !yield(x.Key, x.Val) {
				return false
			}
		}
		return true
	}
}

func Concatenate[K, V any](it1, it2 func(yield func(K, V) bool) bool) func(func(K, V) bool) bool {
	return func(yield func(K, V) bool) bool {
		for k,v := range it1 {
			if !yield(k, v) {
				return false
			}
		}
		for k,v := range it2 {
			if !yield(k, v) {
				return false
			}
		}
		return true
	}
}


// range over function によるrange文を関数として再実装したもの
func MyRange0(f func(yield func()bool)bool, body func()bool) {
	f(body)
}

// range over function によるrange文を関数として再実装したもの
func MyRange1[V any](f func(yield func(V)bool)bool, body func(V) bool) {
	f(body)
}

// range over function によるrange文を関数として再実装したもの
func MyRange2[K, V any](f func(yield func(K, V) bool) bool, body func(K, V) bool) {
	f(body)
}
