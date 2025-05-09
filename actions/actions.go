// Preset actions
package actions

import (
	"fmt"

	"github.com/tsunematsu21/fb"
)

// Preset actions
func FizzBuzz() fb.Action[int] { return func(num int) { fmt.Println("FizzBuzz") } }
func Fizz() fb.Action[int]     { return func(num int) { fmt.Println("Fizz") } }
func Buzz() fb.Action[int]     { return func(num int) { fmt.Println("Buzz") } }
func Pass() fb.Action[int]     { return func(num int) { fmt.Println(num) } }

func Print[T any]() fb.Action[T] {
	return func(a T) {
		fmt.Println(a)
	}
}
