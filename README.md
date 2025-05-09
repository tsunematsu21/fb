# FizzBuzz Generic Library in Go

This is a flexible and reusable FizzBuzz library implemented in Go using generics. It allows you to define custom rules and actions for any type, making it easy to extend the traditional FizzBuzz pattern beyond simple integers.

## Features

- **Generic design:** Works with any type (`T any`).
- **Custom rules:** Define your own matching logic and corresponding actions.
- **Preset rules & actions:** Ready-to-use Fizz, Buzz, FizzBuzz, and fallback utilities for integers.
- **Composable:** Combine multiple rules seamlessly.
- **Utility functions:** Easily handle default behaviors and printing for any type.
- **Minimal and efficient:** Executes the first matching rule per input.

## Presets

The library provides a set of preset actions and rules for the classic integer-based FizzBuzz:

- **Actions:**
  - `actions.FizzBuzz()`
  - `actions.Fizz()`
  - `actions.Buzz()`
  - `actions.Pass()` (prints the number)
  - `actions.Print[T any]()` (generic printer)

- **Rules:**
  - `rules.FizzBuzz()`
  - `rules.Fizz()`
  - `rules.Buzz()`
  - `rules.Pass()` (always matches)
  - `rules.Fallback[T any](action)` (utility fallback rule)

## Installation

```bash
go get github.com/tsunematsu21/fb
```

## Usage
### Basic FizzBuzz
```go:main.go
package main

import (
	"github.com/tsunematsu21/fb"
	"github.com/tsunematsu21/fb/rules"
)

func main() {
	fizzbuzz := fb.New[int](
		rules.FizzBuzz(),
		rules.Fizz(),
		rules.Buzz(),
		rules.Pass(),
	)

	for i := 1; i <= 100; i++ {
		fizzbuzz.Play(i)
	}
}

```

### Custom Rule with Generic Type
```go:main.go
package main

import (
	"fmt"

	"github.com/tsunematsu21/fb"
	"github.com/tsunematsu21/fb/actions"
	"github.com/tsunematsu21/fb/rules"
)

type Item struct {
	Name string
	Qty  int
}

func main() {
	fizzbuzz := fb.New[Item](
		func(item Item) (fb.Action[Item], bool) {
			if item.Qty == 0 {
				return func(item Item) { fmt.Println(item.Name, "is out of stock") }, true
			}
			return nil, false
		},
		rules.Fallback(actions.Print[Item]()),
	)

	fizzbuzz.Play(Item{Name: "Apple", Qty: 5})
	fizzbuzz.Play(Item{Name: "Banana", Qty: 0})
}

```

### Advanced Example: Sekai-no-Nabeatsu
This example reproduces the famous Japanese comedy routine Sekai-no-Nabeatsu, where the comedian counts numbers but acts goofy on numbers that are multiples of 3 or contain the digit 3.
```go:main.go
package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tsunematsu21/fb"
	"github.com/tsunematsu21/fb/rules"
)

func main() {
	multiple := func(num int, a fb.Action[int]) fb.Rule[int] {
		return func(i int) (fb.Action[int], bool) { return a, i%num == 0 }
	}

	contain := func(num int, a fb.Action[int]) fb.Rule[int] {
		return func(i int) (fb.Action[int], bool) {
			return a, strings.Contains(strconv.Itoa(i), strconv.Itoa(num))
		}
	}

	print := func(s string) fb.Action[int] {
		return func(i int) { fmt.Printf("%s<%d\n", s, i) }
	}

	// Sekai-no-Nabeatsu: A Japanese comedy act where the comedian counts numbers
	// but acts goofy on numbers that are multiples of 3 or contain the digit 3.
	nabeatsu := fb.New(
		multiple(3, print("🤪")),
		contain(3, print("🤪")),
		rules.Fallback(print("😀")),
	)

	for i := 1; i <= 100; i++ {
		nabeatsu.Play(i)
	}
}

```

## License
MIT
