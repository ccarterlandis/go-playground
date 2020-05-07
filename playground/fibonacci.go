package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	second_last_fibonacci_num := 0
	last_fibonacci_num := 1
	current_fibonacci_num := 1
	fmt.Println(1)
	return func() int {
		current_fibonacci_num = last_fibonacci_num + second_last_fibonacci_num
		second_last_fibonacci_num = last_fibonacci_num
		last_fibonacci_num = current_fibonacci_num
		return current_fibonacci_num
	}
}

func Fibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
