package main

import (
	"fmt"
	"strings"
)

func main() {
	pass := "abcdefgh"
	hashPass := Hash(pass)

	fmt.Println("hashed:", hashPass)
	fmt.Println(DecyrptHash(hashPass))

}

func Hash(pass string) []int {

	arr := Character()

	hashMap := make(map[string]int, len(arr))

	for index, char := range arr {
		hashMap[char] = index
	}

	splitPass := strings.Split(pass, "")

	var hashPass []int
	for _, char := range splitPass {
		value := hashMap[char]
		hashPass = append(hashPass, value)
	}

	return hashPass
}

func DecyrptHash(pass []int) string {

	arr := Character()

	hashMap := make(map[int]string, len(arr))

	for index, char := range arr {
		hashMap[index] = char
	}
	var decyrptPass string
	for _, char := range pass {
		value := hashMap[char]

		decyrptPass = decyrptPass + value
	}
	return decyrptPass

}

func Character() []string {
	arr := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	arr = append(arr, "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z")
	arr = append(arr, "1", "2", "3", "4", "5", "6", "7", "8", "9", "0")

	return arr
}
