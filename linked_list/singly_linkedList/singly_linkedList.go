package singly_linkedList

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: &Node{}}
}

func (l *LinkedList) Prepend(val *Node) {
	tmp := l.head
	l.head = val
	l.head.next = tmp
	l.length++
}

func (l *LinkedList) UpdateByVal(fromVal, toVal int) {
	toUpdate, i := l.head, l.length

	for i > 0 {
		if toUpdate.data == fromVal {
			toUpdate.data = toVal
		}
		toUpdate = toUpdate.next
		i--
	}
}

func (l *LinkedList) DeleteByVal(val int) {
	if l.length == 0 {
		return
	}

	if l.head.data == val {
		l.head = l.head.next
		l.length--
		return
	}

	nodeToDelete := l.head
	for nodeToDelete.next.data != val {
		if nodeToDelete.next.next == nil {
			return
		}
		nodeToDelete = nodeToDelete.next
	}
	nodeToDelete.next = nodeToDelete.next.next
	l.length--
}

func (l LinkedList) Print() {
	toPrint := l.head
	for l.length > 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l LinkedList) SearchByVal(val int) bool {
	toSearch := l.head
	for l.length > 0 {
		if toSearch.data == val {
			return true
		}
		toSearch = toSearch.next
		l.length--
	}
	return false
}

/*
func main() {
	ll := NewLinkedList()
	ll.Prepend(&Node{data: 56})
	ll.Prepend(&Node{data: 31})
	ll.Prepend(&Node{data: 43})
	ll.Prepend(&Node{data: 37})
	ll.Prepend(&Node{data: 2})
	ll.Print()
	ll.DeleteByVal(31)
	ll.Print()
	fmt.Println(ll.SearchByVal(43))
	ll.UpdateByVal(56, 89)
	ll.Print()
}
*/
