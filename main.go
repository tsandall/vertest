package main

import "fmt"
import "github.com/tsandall/vertest/version"

func fib(x int) int {
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	return fib(x-1) + fib(x-2)
}

func main() {
	fmt.Printf("Goodbye Cruel World (%v-%v)\n", version.Version, version.Vcs)
	fmt.Println(fib(10))
}
