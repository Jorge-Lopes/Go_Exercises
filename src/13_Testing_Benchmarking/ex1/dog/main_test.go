package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	n := Years(2)
	if n != 14 {
		t.Error("Expected 14, got:", n)
	}
}

func TestYearsTwo(t *testing.T) {
	n := YearsTwo(2)
	if n != 14 {
		t.Error("Expected 14, got:", n)
	}
}

func ExampleYears() {
	fmt.Println(Years(3))
	// Output:
	// 21
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(3))
	// Output:
	// 21
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(3)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(3)
	}
}
