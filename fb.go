/*
Package fb provides a generic and extensible FizzBuzz implementation using Go generics.

You can define your own matching rules and actions for any type (not just integers).
Preset rules and actions for classic integer-based FizzBuzz are also provided.
*/
package fb

// Action is a function that takes a value of type T and performs an action (e.g., print).
type Action[T any] func(T)

// Rule is a function that takes a value of type T and returns an Action (if it matches) and a bool indicating a match.
//
// If the rule matches, the corresponding Action will be executed by FizzBuzz.
type Rule[T any] func(T) (Action[T], bool)

// FizzBuzz holds a set of rules and applies them to values using the Play method.
//
// The first matching rule will be executed.
type FizzBuzz[T any] struct {
	Rules []Rule[T]
}

// Play applies the FizzBuzz rules to the given value.
//
// It iterates over the rules in order and executes the first rule that matches.
// If a matching rule returns a non-nil Action, the Action is executed with the value.
// If no rule matches or the Action is nil, Play does nothing.
//
// This method is safe even if a Rule returns a nil Action.
func (f *FizzBuzz[T]) Play(val T) {
	for _, rule := range f.Rules {
		if act, ok := rule(val); ok {
			if act != nil {
				act(val)
			}
			return
		}
	}
}

// New creates a new FizzBuzz instance with the provided rules.
//
// Example usage:
//
//	fb.New[int](rules.FizzBuzz(), rules.Fizz(), rules.Buzz(), rules.Pass())
func New[T any](rules ...Rule[T]) *FizzBuzz[T] {
	fb := &FizzBuzz[T]{Rules: rules}
	return fb
}
