package fb

type Action[T any] func(T)

type Rule[T any] func(T) (Action[T], bool)

type FizzBuzz[T any] struct {
	Rules []Rule[T]
}

func (f *FizzBuzz[T]) Play(val T) {
	for _, rule := range f.Rules {
		if act, ok := rule(val); ok {
			act(val)
			return
		}
	}
}

func New[T any](rules ...Rule[T]) *FizzBuzz[T] {
	fb := &FizzBuzz[T]{Rules: rules}
	return fb
}
