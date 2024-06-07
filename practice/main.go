package main

import "fmt"

func fib(n int) int {
	if n == 1 {
		return 1
	}
	return fib(n-1) * n
}

func main() {
	fmt.Printf("%d", fib(3))
}
