/*
Package fb implements a generic, extensible FizzBuzz using Go generics.

It supports custom rules and actions for any type, with presets for classic integer-based FizzBuzz.
*/
package fb

// Action is a function that processes a value of type T.
type Action[T any] func(T)

// Rule checks a value and returns an Action (if matched) and a bool indicating a match.
type Rule[T any] func(T) (Action[T], bool)

// FizzBuzz applies rules to values and runs the first matching Action.
type FizzBuzz[T any] struct {
	Rules []Rule[T]
}

// Play applies rules to val, running the first matching Action if non-nil.
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

// New returns a FizzBuzz instance with the given rules.
//
// Example:
//
//	fb.New[int](rules.FizzBuzz(), rules.Fizz(), rules.Buzz(), rules.Pass())
func New[T any](rules ...Rule[T]) *FizzBuzz[T] {
	return &FizzBuzz[T]{Rules: rules}
}
