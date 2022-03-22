package main

import (
	"fmt"
	"math"
	"strconv"
)

type Stack []int

func (st *Stack) IsEmpty() bool {
	return len(*st) == 0
}

func (st *Stack) Push(data int) {
	*st = append(*st, data)
	fmt.Println("After push ", st)
}
func (st *Stack) Pop() bool {
	if st.IsEmpty() {
		return false
	} else {
		index := len(*st) - 1
		*st = (*st)[:index]
		return true
	}
}
func (st *Stack) Top() int {
	if st.IsEmpty() {
		return 0
	} else {
		index := len(*st) - 1
		element := (*st)[index]
		fmt.Println("After Pops", st)
		return element
	}
}
func EvolationPostfix(postfix string) int {
	var intStack Stack
	for _, char := range postfix {
		opchar := string(char)
		if opchar >= "0" && opchar <= "9" {
			i1, _ := strconv.Atoi(opchar)
			intStack.Push(i1)
		} else {
			opr1 := intStack.Top()
			intStack.Pop()
			opr2 := intStack.Top()
			intStack.Pop()

			switch char {
			case '^':
				x := math.Pow(float64(opr2), float64(opr1))
				intStack.Push(int(x))
			case '+':
				intStack.Push(opr2 + opr1)
			case '-':
				intStack.Push(opr2 - opr1)
			case '*':
				intStack.Push(opr2 * opr1)
			case '/':
				intStack.Push(opr2 / opr1)
			}

		}
	}
	return intStack.Top()
}
func main() {
	postfix := "2323^5-212*+^*+4-"
	evaluationReslt := EvolationPostfix(postfix)
	fmt.Printf("evaluation of %s is %d", postfix, evaluationReslt)
}
