package go61405_test

import "fmt"

func ExampleRangeInteger() {
	type MyInt int
	for i := range MyInt(3) {
		fmt.Printf("type=%T, value=%d\n", i, i)
	}
	// Output:
	// type=go61405_test.MyInt, value=0
	// type=go61405_test.MyInt, value=1
	// type=go61405_test.MyInt, value=2
}
