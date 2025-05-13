package match

import "testing"

func TestOneAddTwo(t *testing.T) {
	num1 := 1
	num2 := 2
	expected := 3

	actual := Add(num1, num2)

	if actual != expected {
		t.Errorf("Add(%d, %d) = %d; expected %d", num1, num2, actual, expected)
	}
}
