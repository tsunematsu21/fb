package fb_test

import (
	"fmt"
	"testing"

	"github.com/tsunematsu21/fb"
	"github.com/tsunematsu21/fb/rules"
)

func TestFizzBuzz_Play_FirstMatchOnly(t *testing.T) {
	var result []string

	// Define some rules that append to result slice.
	rule1 := func(i int) (fb.Action[int], bool) {
		return func(i int) { result = append(result, "Rule1") }, i%2 == 0
	}
	rule2 := func(i int) (fb.Action[int], bool) {
		return func(i int) { result = append(result, "Rule2") }, i%3 == 0
	}
	fallback := func(i int) (fb.Action[int], bool) {
		return func(i int) { result = append(result, "Fallback") }, true
	}

	fizzbuzz := fb.New(rule1, rule2, fallback)

	tests := []struct {
		input    int
		expected string
	}{
		{2, "Rule1"},
		{3, "Rule2"},
		{5, "Fallback"},
		{6, "Rule1"}, // Rule1 should win over Rule2 if both match
	}

	for _, tt := range tests {
		result = []string{}
		fizzbuzz.Play(tt.input)
		if len(result) != 1 || result[0] != tt.expected {
			t.Errorf("input %d: expected %s, got %v", tt.input, tt.expected, result)
		}
	}
}

func TestGenericType(t *testing.T) {
	type Item struct {
		Name string
		Qty  int
	}
	var result []string

	rule := func(item Item) (fb.Action[Item], bool) {
		return func(item Item) { result = append(result, fmt.Sprintf("%s: %d", item.Name, item.Qty)) }, true
	}

	fizzbuzz := fb.New(rule)

	item := Item{Name: "Apple", Qty: 10}
	fizzbuzz.Play(item)

	if len(result) != 1 || result[0] != "Apple: 10" {
		t.Errorf("unexpected result: %v", result)
	}
}

func ExampleFizzBuzz_Play() {
	fizzbuzz := fb.New(
		rules.FizzBuzz(),
		rules.Fizz(),
		rules.Buzz(),
		rules.Pass(),
	)

	for i := 1; i <= 15; i++ {
		fizzbuzz.Play(i)
	}

	// Output:
	// 1
	// 2
	// Fizz
	// 4
	// Buzz
	// Fizz
	// 7
	// 8
	// Fizz
	// Buzz
	// 11
	// Fizz
	// 13
	// 14
	// FizzBuzz
}

func TestNilFizzBuzz(t *testing.T) {
	var fb *fb.FizzBuzz[int]
	defer noPanic(t, "Play on nil FizzBuzz")
	fb.Play(1) // should do nothing, not panic
}

func TestNilRule(t *testing.T) {
	fb := fb.New[int](
		nil,
		func(i int) (fb.Action[int], bool) {
			if i%2 == 0 {
				return func(i int) {}, true
			}
			return nil, false
		},
	)
	defer noPanic(t, "Play with nil Rule in rules")
	fb.Play(2)
}

func TestNilAction(t *testing.T) {
	fb := fb.New[int](
		func(i int) (fb.Action[int], bool) {
			return nil, true // matching rule, but action is nil
		},
	)
	defer noPanic(t, "Play with Rule returning nil Action")
	fb.Play(1)
}

func TestNilRuleAndAction(t *testing.T) {
	fb := fb.New[int](
		nil,
		func(i int) (fb.Action[int], bool) {
			return nil, true
		},
	)
	defer noPanic(t, "Play with nil Rule and nil Action")
	fb.Play(1)
}

// Helper: fails the test if panic occurs
func noPanic(t *testing.T, name string) {
	if r := recover(); r != nil {
		t.Fatalf("%s: unexpected panic: %v", name, r)
	}
}
