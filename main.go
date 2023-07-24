package main

import "fmt"

func main() {
	for x, y := range F {
		fmt.Println("range F", x, y)
	}
}

func F(yield func(id int, name string) bool) bool {
	for _, x := range []struct {
		id   int
		name string
	}{
		{0, "a"},
		{1, "b"},
		{2, "c"},
		{3, "d"},
	} {
		fmt.Println("Before invoke yield", x.id, x.name)
		if result := yield(x.id, x.name); !result {
			return false
		}
		fmt.Println("After invoke yield", x.id, x.name)
	}
	return false
}
