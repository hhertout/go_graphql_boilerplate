package loop_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/hhertout/twirp_auth/lib/loop"
)

func TestMap_IntToString(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []string{"1", "2", "3", "4", "5"}

	result := loop.Map(input, func(value int) string {
		return strconv.Itoa(value)
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

func TestMap_StringToUpper(t *testing.T) {
	input := []string{"hello", "world"}
	expected := []string{"HELLO", "WORLD"}

	result := loop.Map(input, func(value string) string {
		return strings.ToUpper(value)
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

func TestMap_EmptyInput(t *testing.T) {
	input := []int{}
	expected := []int{}

	result := loop.Map(input, func(value int) int {
		return value * 2
	})

	if len(result) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(result))
	}

	if len(result) != 0 {
		t.Errorf("expected empty result, got %v", result)
	}
}

func TestMap_StructToStruct(t *testing.T) {
	type Input struct {
		A int
		B string
	}
	type Output struct {
		C int
		D string
	}

	input := []Input{
		{A: 1, B: "one"},
		{A: 2, B: "two"},
	}
	expected := []Output{
		{C: 1, D: "ONE"},
		{C: 2, D: "TWO"},
	}

	result := loop.Map(input, func(value Input) Output {
		return Output{
			C: value.A,
			D: strings.ToUpper(value.B),
		}
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

func TestMap_NilInput(t *testing.T) {
	var input []int
	expected := []int{}

	result := loop.Map(input, func(value int) int {
		return value * 2
	})

	if result == nil {
		t.Fatalf("expected non-nil result, got nil")
	}

	if len(result) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(result))
	}
}

func TestMap_DifferentTypes(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []float64{1.1, 2.1, 3.1}

	result := loop.Map(input, func(value int) float64 {
		return float64(value) + 0.1
	})

	if len(result) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(result))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("expected %f at index %d, got %f", expected[i], i, v)
		}
	}
}
