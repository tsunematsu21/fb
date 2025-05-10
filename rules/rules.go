// Package rules provides preset Rule implementations for fb.
//
// Includes standard FizzBuzz rules and utilities like Pass and Fallback.
//
// Example:
//
//	rule := rules.Fizz()
//	if act, ok := rule(3); ok {
//	    act(3)
//	}
package rules

import (
	"github.com/tsunematsu21/fb"
	"github.com/tsunematsu21/fb/actions"
)

// FizzBuzz returns a Rule that matches numbers divisible by 15.
func FizzBuzz() fb.Rule[int] {
	return func(i int) (fb.Action[int], bool) { return actions.FizzBuzz(), i%15 == 0 }
}

// Fizz returns a Rule that matches numbers divisible by 3.
func Fizz() fb.Rule[int] {
	return func(i int) (fb.Action[int], bool) { return actions.Fizz(), i%3 == 0 }
}

// Buzz returns a Rule that matches numbers divisible by 5.
func Buzz() fb.Rule[int] {
	return func(i int) (fb.Action[int], bool) { return actions.Buzz(), i%5 == 0 }
}

// Pass returns a Rule that always matches (used as fallback).
func Pass() fb.Rule[int] {
	return func(i int) (fb.Action[int], bool) { return actions.Pass(), true }
}

// Fallback returns a Rule that always matches and runs the provided Action.
//
// This is a generic utility fallback rule for any type.
func Fallback[T any](a fb.Action[T]) fb.Rule[T] {
	return func(t T) (fb.Action[T], bool) { return a, true }
}

// Match creates a rule that applies the given predicate to the value and returns an Action if true.
func Match[T any](predicate func(T) bool, action fb.Action[T]) fb.Rule[T] {
	return func(value T) (fb.Action[T], bool) {
		if predicate(value) {
			return action, true
		}
		return nil, false
	}
}
