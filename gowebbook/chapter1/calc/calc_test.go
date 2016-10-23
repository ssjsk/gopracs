package calc

import "testing"

func TestAdd(t *testing.T){
	var result int
	result = Add(15, -5)
	if result != 10 {
		t.Error("Expected 10, got ", result)
	}
}
func TestSubtract(t *testing.T){
	var result int
	result = Subtract(15, 10)

	if result != 5 {
		t.Error("Expected 5, got ", result)
	}
}