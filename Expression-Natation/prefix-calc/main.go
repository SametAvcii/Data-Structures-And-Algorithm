package main

import (
	"fmt"
	"math"
	"strconv"
)

type Stack []int
type Byte byte

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
func EvolationInfix(infix string) int {
	var intStack Stack
	for j := len(infix) - 1; j >= 0; j-- {
		char := (infix)[j]
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
				x := math.Pow(float64(opr1), float64(opr2))
				intStack.Push(int(x))
			case '+':
				intStack.Push(opr1 + opr2)
			case '-':
				intStack.Push(opr1 - opr2)
			case '*':
				intStack.Push(opr1 * opr2)
			case '/':
				intStack.Push(opr1 / opr2)
			}

		}
	}
	return intStack.Top()
}
func main() {
	infix := "++*23^-5721"
	evaluationReslt := EvolationInfix(infix)
	fmt.Printf("evaluation of %s is \n result %d  ", infix, evaluationReslt)
}
