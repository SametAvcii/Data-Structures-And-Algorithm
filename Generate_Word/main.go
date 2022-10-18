package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func GenerateCombinations(words string, length int) <-chan string {
	c := make(chan string)

	go func(c chan string) {
		defer close(c)

		CreateWord(c, "", words, length)
	}(c)

	return c
}

func CreateWord(c chan string, combo string, words string, length int) {

	if length <= 0 {
		return
	}

	var newCombo string
	for _, letter := range words {
		newCombo = combo + string(letter)
		c <- newCombo
		CreateWord(c, newCombo, words, length-1)
	}
}

func main() {

	arrUnique := createUniqueArray("123", "456", "789")

	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	var word string

	for _, value := range arrUnique {
		word += value
	}
	for combination := range GenerateCombinations(word, 5) {
		defer f.Close()

		if len(combination) != 5 {
			continue
		}
		_, err2 := f.WriteString(combination + "\n")
		if err2 != nil {
			log.Fatal(err2)
		}
	}

	fmt.Println("done")

}
func createUniqueArray(word1, word2, word3 string) []string {

	arr1 := strings.Split(word1, "")
	arr2 := strings.Split(word2, "")
	arr3 := strings.Split(word3, "")

	arrUnique := append(arr1, arr2...)
	arrUnique = append(arrUnique, arr3...)

	arrUnique = removeDuplicates(arrUnique)

	return arrUnique
}

func removeDuplicates(strList []string) []string {
	list := []string{}
	for _, item := range strList {
		if contains(list, item) == false {
			list = append(list, item)
		}
	}
	return list
}

func contains(list []string, value string) bool {
	for _, duplicate := range list {
		if duplicate == value {
			return true
		}
	}
	return false
}
