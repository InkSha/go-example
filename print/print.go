package example

import "fmt"

func ExamplePrint(word string, n int) {

	for i := 0; i < n; i++ {
		fmt.Println(word)
	}
}
