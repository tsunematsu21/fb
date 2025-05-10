// Package actions provides preset Action implementations for fb.
//
// Includes "Fizz", "Buzz", "FizzBuzz", printing numbers, and a generic Print for any type.
//
// Example:
//
//	actions.Fizz()(3) // Output: Fizz
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

// Print returns an Action that prints a value of type T.
func Print[T any]() fb.Action[T] {
	return func(a T) {
		fmt.Println(a)
	}
}

// Noop returns a nil Action for type T.
func Noop[T any]() fb.Action[T] {
	return nil
}
