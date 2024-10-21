package loop_test

import (
	"testing"

	"github.com/hhertout/twirp_auth/lib/loop"
)

func TestFilter_Ints(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4}

	result := loop.Filter(input, func(value int) bool {
		return value%2 == 0
	})

	if len(result) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(result))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("expected %d at index %d, got %d", expected[i], i, v)
		}
	}
}

func TestFilter_Strings(t *testing.T) {
	input := []string{"apple", "banana", "cherry", "date"}
	expected := []string{"apple"}

	result := loop.Filter(input, func(value string) bool {
		return len(value) == 5
	})

	if len(result) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(result))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("expected %s at index %d, got %s", expected[i], i, v)
		}
	}
}

func TestFilter_EmptyInput(t *testing.T) {
	input := []int{}
	expected := []int{}

	result := loop.Filter(input, func(value int) bool {
		return value%2 == 0
	})

	if len(result) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(result))
	}

	if len(result) != 0 {
		t.Errorf("expected empty result, got %v", result)
	}
}

func TestFilter_Structs(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	input := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}
	expected := []Person{
		{Name: "Alice", Age: 30},
		{Name: "Charlie", Age: 35},
	}

	result := loop.Filter(input, func(value Person) bool {
		return value.Age >= 30
	})

	if len(result) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(result))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("expected %v at index %d, got %v", expected[i], i, v)
		}
	}
}

func TestFilter_NilCallback(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("expected panic, got none")
		}
	}()

	input := []int{1, 2, 3}
	loop.Filter(input, nil)
}

func TestFilter_NilInput(t *testing.T) {
	var input []int
	expected := []int{}

	result := loop.Filter(input, func(value int) bool {
		return value%2 == 0
	})

	if result == nil {
		t.Fatalf("expected non-nil result, got nil")
	}

	if len(result) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(result))
	}
}
