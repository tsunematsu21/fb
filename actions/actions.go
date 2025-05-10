// Package actions provides preset Action implementations for use with the fb library.
//
// These include common actions for integer-based FizzBuzz scenarios, such as printing "Fizz", "Buzz",
// "FizzBuzz", or simply the number itself. It also provides a generic Print action for any type.
//
// Example usage:
//
//	fizzAction := actions.Fizz()
//	fizzAction(3) // Output: Fizz
package actions

import (
	"fmt"

	"github.com/tsunematsu21/fb"
)

// FizzBuzz returns an Action that prints "FizzBuzz".
func FizzBuzz() fb.Action[int] { return func(num int) { fmt.Println("FizzBuzz") } }

// Fizz returns an Action that prints "Fizz".
func Fizz() fb.Action[int] { return func(num int) { fmt.Println("Fizz") } }

// Buzz returns an Action that prints "Buzz".
func Buzz() fb.Action[int] { return func(num int) { fmt.Println("Buzz") } }

// Pass returns an Action that prints the number itself.
func Pass() fb.Action[int] { return func(num int) { fmt.Println(num) } }

// Print returns an Action that prints any type T.
//
// This is useful for custom types.
func Print[T any]() fb.Action[T] {
	return func(a T) {
		fmt.Println(a)
	}
}
