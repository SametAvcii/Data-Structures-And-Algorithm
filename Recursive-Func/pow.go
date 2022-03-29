package main

import "fmt"

func Pow(a, n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return a
	} else {
		pow := Pow(a, n-1)
		return pow * a
	}
}
func PowLog(a, n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return a
	} else {
		if n%2 == 0 {
			pow := PowLog(a, n/2)
			return pow * pow
		} else {
			pow := PowLog(a, (n-1)/2)
			return a * pow * pow
		}
	}
}

func main() {
	fmt.Println(Pow(2, 3))
	fmt.Println(PowLog(2, 3))
}
