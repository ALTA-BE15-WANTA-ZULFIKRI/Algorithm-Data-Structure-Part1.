package main

import (
	"fmt"
)
// Fibonacci adalah sebuah fungsi yang mengembalikan nilai dari deret Fibonacci
func Fibonacci(number int) int {
	if number == 0 {
		return 0
	} else if number == 1 {
		return 1
	} else {
		return Fibonacci(number-1) + Fibonacci(number-2)
	}
}

func main() {

	fmt.Println(Fibonacci(0))  // 0
	fmt.Println(Fibonacci(2))  // 1
	fmt.Println(Fibonacci(9))  // 34
	fmt.Println(Fibonacci(10)) // 55
	fmt.Println(Fibonacci(12)) // 144

}
