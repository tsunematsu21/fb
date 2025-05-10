/*
Package fb implements a generic, extensible FizzBuzz using Go generics.

It supports custom rules and actions for any type, with presets for classic integer-based FizzBuzz.
*/
package fb

// Action is a function that processes a value of type T.
type Action[T any] func(T)

// Rule returns an Action and true if the value matches; otherwise, false.
type Rule[T any] func(T) (Action[T], bool)

// FizzBuzz holds a list of rules and applies them to values.
type FizzBuzz[T any] struct {
	rules []Rule[T]
}

// Play applies rules to val, running the first matching Action if non-nil.
func (f *FizzBuzz[T]) Play(val T) {
	if f == nil {
		return
	}

	for _, rule := range f.rules {
		if rule == nil {
			continue
		}
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
	return &FizzBuzz[T]{rules: rules}
}
