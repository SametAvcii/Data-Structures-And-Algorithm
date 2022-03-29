package main

import "fmt"

func RecursiveFunc(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * RecursiveFunc(n-1)
	}
}
func main() {
	fmt.Println(IteratifFunc(5))
	fmt.Println(Sum(6))
}

func IteratifFunc(n int) int {
	res := 1
	for i := n; i > 0; i-- {
		res *= i
	}
	return res
}

func Sum(n int) int {
	if n == 1 {
		return 1
	} else {
		sum := 0
		sum = Sum(n - 1)
		return sum + n
	}

}
