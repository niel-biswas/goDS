package doubly_LinkedList

import "fmt"

type Node struct {
	data     int
	next     *Node
	previous *Node
}

type DoublyLinkedList struct {
	head   *Node
	length int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{head: &Node{}}
}

func (ll *DoublyLinkedList) Prepend(val *Node) {
	if ll.length == 0 {
		ll.head = val
		ll.length++
		return
	}
	val.next = ll.head
	ll.head.previous = val
	ll.head = val
	ll.length++
}

func (ll *DoublyLinkedList) Append(val *Node) {
	if ll.length == 0 {
		ll.head = val
		ll.length++
		return
	}
	nodeToAppend := ll.head
	for nodeToAppend.next != nil {
		nodeToAppend = nodeToAppend.next
	}
	nodeToAppend.next = val
	val.previous = nodeToAppend
	ll.length++
}

func (ll *DoublyLinkedList) DeleteByVal(val int) {
	if ll.length == 0 {
		return
	}

	if ll.head.data == val {
		ll.head = ll.head.next
		ll.length--
	}

	previousNodeToDelete, previousNode := ll.head, &Node{}
	for previousNodeToDelete.next.data != val {
		previousNode = previousNodeToDelete
		if previousNodeToDelete.next.next == nil {
			return
		}
		previousNodeToDelete = previousNodeToDelete.next
	}
	previousNodeToDelete.next = previousNodeToDelete.next.next
	previousNodeToDelete.previous = previousNode
	ll.length--
}

func (ll *DoublyLinkedList) SearchByVal(val int) bool {
	if ll.length == 0 {
		return false
	}
	for node := ll.head; node.next != nil; node = node.next {
		if node.data == val {
			return true
		}
	}
	return false
}

func (ll *DoublyLinkedList) UpdateByVal(fromVal, toVal int) {
	if ll.length == 0 {
		return
	}
	for node := ll.head; node.next != nil; node = node.next {
		if node.data == fromVal {
			node.data = toVal
		}
	}
}

func (ll DoublyLinkedList) Print() {
	nodeToPrint := ll.head
	for ll.length > 0 {
		fmt.Printf("%d ", nodeToPrint.data)
		nodeToPrint = nodeToPrint.next
		ll.length--
	}
	fmt.Printf("\n")
}

/*func main() {
	ll := NewDoublyLinkedList()
	ll.Append(&Node{data: 56})
	ll.Append(&Node{data: 31})
	ll.Append(&Node{data: 43})
	ll.Append(&Node{data: 37})
	ll.Append(&Node{data: 12})
	ll.Prepend(&Node{data: 2})
	ll.Print()
	ll.DeleteByVal(31)
	ll.Print()
	fmt.Println(ll.SearchByVal(43))
	ll.UpdateByVal(56, 89)
	ll.Print()
}*/
