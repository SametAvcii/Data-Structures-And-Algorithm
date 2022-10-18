package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	pass := "Merhaba YM-LAB-3"

	hashPass := Hash(pass)
	decPass := DecyrptHash(hashPass)

	fmt.Println("hashed:", hashPass)

	fmt.Println("decPass:",decPass)

}

func Hash(pass string) string {

	passArr := strings.Split(pass, "")

	var arr []string
	for _, passChar := range passArr {

		value := ""
		asciiValue := []byte(passChar)
		sumDigit, count := findDigitSum(int(asciiValue[0]))

		intAscii := int(asciiValue[0])

		digits := strings.Split(strconv.Itoa(intAscii), "")
		firstDigit := digits[0]
		if len(digits)==3{
			firstDigit+=digits[1]
		}

		value += firstDigit + strconv.Itoa(sumDigit) + strconv.Itoa(count)
		arr = append(arr, value)

	}


	hashValue := ""

	for _, val := range arr {

		hashValue += "180" + val
	}
	fmt.Println(hashValue)
	return hashValue
}
func findDigitSum(num int) (int, int) {
	res := 0
	count := 0
	for num > 0 {
		res += num % 10
		num /= 10
		count++
	}
	return res, count
}

func DecyrptHash(pass string) string {

	splitArr := strings.Split(pass, "180")
	splitArr=splitArr[1:]


	var asciiArr []int

	for _, val := range splitArr {
	var ascVal string
		
		digit:= strings.Split(val, "")

		ascVal += digit[0]

		lastIndis := len(digit) - 1
		if digit[(lastIndis)] == "3" {

			ascVal += digit[1]
			digit0, _ := strconv.Atoi(digit[0]) 
			digit1, _ := strconv.Atoi(digit[1])

			digit = digit[2 : len(digit)-1]
			var total string

			for _, tot := range digit {
				total += tot
			}

			tot, _ := strconv.Atoi(total)
			thirdDigit := tot - (digit0 + digit1)
			thirdDigitStr := strconv.Itoa(thirdDigit)
			ascVal += thirdDigitStr
		} else if digit[(lastIndis)] == "2" {
			digit0, _ := strconv.Atoi(digit[0])

			digit = digit[1 : len(digit)-1]
			var total string

			for _, tot := range digit {
				total += tot
			}

			tot, _ := strconv.Atoi(total)
			secDigit := tot - digit0

			secDigitStr := strconv.Itoa(secDigit)

			ascVal += secDigitStr
		} else {
			ascVal += digit[0]
		}

		indis, _ := strconv.Atoi(ascVal)
		asciiArr = append(asciiArr, indis)
	}
	
	
	var decHash string
	for _, ascii := range asciiArr {

		decHash+=string(ascii)
	}

	return decHash

}

