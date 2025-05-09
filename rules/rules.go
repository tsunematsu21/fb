package rules

import (
	"github.com/tsunematsu21/fb"
	"github.com/tsunematsu21/fb/actions"
)

// Preset rules
func FizzBuzz() fb.Rule[int] {
	return func(i int) (fb.Action[int], bool) { return actions.FizzBuzz(), i%15 == 0 }
}
func Fizz() fb.Rule[int] {
	return func(i int) (fb.Action[int], bool) { return actions.Fizz(), i%3 == 0 }
}
func Buzz() fb.Rule[int] {
	return func(i int) (fb.Action[int], bool) { return actions.Buzz(), i%5 == 0 }
}
func Pass() fb.Rule[int] {
	return func(i int) (fb.Action[int], bool) { return actions.Pass(), true }
}

// Utility rules
func Fallback[T any](a fb.Action[T]) fb.Rule[T] {
	return func(t T) (fb.Action[T], bool) { return a, true }
}
