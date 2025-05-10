package rules_test

import (
	"testing"

	"github.com/tsunematsu21/fb/rules"
)

func TestMatch(t *testing.T) {
	// Action that we expect to be executed
	action := func(num int) { /* do something */ }

	// Test case where the predicate matches
	rule := rules.Match(func(i int) bool { return i%2 == 0 }, action)
	act, matched := rule(2)
	if !matched {
		t.Errorf("Expected match, but got no match")
	}
	if act == nil {
		t.Errorf("Expected action, but got nil")
	}

	// Test case where the predicate does not match
	act, matched = rule(3)
	if matched {
		t.Errorf("Expected no match, but got match")
	}
	if act != nil {
		t.Errorf("Expected nil action, but got %v", act)
	}
}
