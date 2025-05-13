package example

import "testing"

func TestPrintThreeHelloWorld(t *testing.T) {
	ExamplePrint("hello world", 3)
	// Output:
	// hello world
	// hello world
	// hello world
}
