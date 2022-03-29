package main

import "fmt"

func Fibonacci(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 || n == 3 {
		return 1
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}
func FibonacciIterative(n int) int {

	if n == 1 {
		return 0
	} else if n == 2 || n == 3 {
		return 1
	} else {
		first, second, sum := 1, 1, 0
		for i := 3; i < n; i++ {
			sum = first + second
			first = second
			second = sum
		}
		return sum
	}

}
func main() {
	fmt.Println(Fibonacci(50))
	fmt.Println(FibonacciIterative(50))
	//Kodun notasyonunu ve hangi amaçla yazıldığına
	//https://medium.com/@SametAvcii/recursive-yap%C4%B1lar-593cf39484e7 linkten ulaşabilirsiniz.
}
