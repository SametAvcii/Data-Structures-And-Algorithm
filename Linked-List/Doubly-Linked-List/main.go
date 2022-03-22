package main

import (
	"fmt"
)

type ListNode struct {
	data int
	prev *ListNode
	next *ListNode
}
type LinkedList struct {
	Head *ListNode
	len  int
}

func (l *LinkedList) AddNodeToFront(data int) {
	n := ListNode{}
	n.data = data

	if l.len == 0 {
		l.Head = &n
		l.len++
		return
	} else {

		n.next = l.Head
		l.Head.prev = &n
		l.Head = &n
		l.len++
		return
	}

}
func (l *LinkedList) AddNodeToBack(data int) {

	if l.len == 0 {
		n := ListNode{}
		n.data = data
		l.Head = &n
		l.len++

	} else {
		last := ListNode{data: data}
		current := l.Head
		for current.next != nil {
			current = current.next
		}
		current.next = &last
		last.prev = current
		l.len++
		return

	}

}

func (l *LinkedList) AddNodeToBetween(data, num int) {
	n := ListNode{}

	if l.len == 0 {
		n.data = data
		l.Head = &n
		l.len++
	} else {
		current := l.Head
		prev := current

		for i := 0; i < num-1; i++ {
			prev = current
			current = current.next
		}

		n.data = data
		prev.next = &n
		n.prev = prev
		n.next = current
		current.prev = &n
		l.len++
		return
	}
}
func (l *LinkedList) RemoveFromFront() {
	if l.len == 0 {
		fmt.Println("Don't have Node")
	}
	l.Head = l.Head.next
	l.Head.next.prev = nil
	l.len--
}
func (l *LinkedList) RemoveFromBack() {
	if l.len == 0 {
		fmt.Println("Don't have Node")
	}
	var prev *ListNode
	current := l.Head
	for current.next != nil {
		prev = current
		current = current.next
	}
	if prev != nil {
		prev.next = nil
		current.prev = nil
	} else {
		l.Head = nil
	}
	l.len--
	return
}
func (l *LinkedList) RemoveFromBetween() {
	if l.len == 0 {
		fmt.Println("Don't have Node")
	}

	var prev *ListNode
	current := l.Head
	for i := 0; i < 2; i++ {
		prev = current
		current = current.next
	}
	prev.next = current.next
	l.len--
	return

}
func (l *LinkedList) RemoveFromData(data int) {

	if l.Head == nil {
		fmt.Println("List is Empty")
	} else {
		var prev *ListNode
		current := l.Head

		for i := 0; i < l.len; i++ {

			prev = current
			if current.data != data {
				current = current.next
			} else {
				if current.next != nil {
					prev.next = current.next
				} else {
					prev.next = nil
				}

				fmt.Printf("Remove the %d", data)
				fmt.Println()
				l.len--
				break
			}

		}
		fmt.Printf("%d not found in the list", data)
		fmt.Println()

	}
}

func (l *LinkedList) RemoveAll() {
	if l.Head == nil {
		fmt.Println("List is Empty")
	}
	var prev *ListNode
	current := l.Head
	for current.next != nil {
		current.prev = nil
		prev = current
		current = current.next
		prev.next = nil
	}
	l.Head = nil
	l.len = 0
	return
}
func (l *LinkedList) ListOfAll() {
	if l.Head == nil {
		fmt.Println("List is Empty")
	}
	current := l.Head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}

}
func (l *LinkedList) SearchInList(data int) {

	if l.Head == nil {
		fmt.Println("List is empty")
	} else {
		count := 1
		current := l.Head

		for i := 0; i < l.len; i++ {
			if current.data == data {
				fmt.Printf("Your Data %d your data count : %d", current.data, count)
				fmt.Println()
				break
			} else {
				current = current.next
				count++
			}
		}
		fmt.Println(current.data)

		if current.data != data {
			fmt.Printf("%d not found in list", data)
			fmt.Println()
		}

	}
	return
}
func InitList() *LinkedList {
	return &LinkedList{}
}
func main() {

	linkedList := InitList()
	linkedList.AddNodeToFront(50)
	linkedList.AddNodeToFront(50)
	linkedList.AddNodeToFront(50)
	linkedList.AddNodeToFront(200)
	fmt.Println(linkedList.len)
	linkedList.AddNodeToBetween(100, 2)
	linkedList.RemoveFromData(200)
	fmt.Println(linkedList.len)

	linkedList.ListOfAll()
	linkedList.SearchInList(200)

}
