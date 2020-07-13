package linked

import "strings"

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	Lenght int
	head   *Node
}

func CreateLL() *LinkedList {
	ll := LinkedList{Lenght: 0, head: nil}
	return &ll
}

func (ll *LinkedList) Add(data string) {
	newNode := Node{data: data, next: nil}

	// Empty list
	if ll.head == nil {
		ll.head = &newNode
		ll.Lenght++
		return
	}

	tail := ll.head
	for tail.next != nil {
		tail = tail.next
	}
	tail.next = &newNode
	ll.Lenght++
}

func (ll LinkedList) String() string {
	if ll.head == nil {
		return "The list empty"
	}

	content := make([]string, 0, ll.Lenght)
	node := ll.head
	for node.next != nil {
		content = append(content, node.data)
		node = node.next
	}
	content = append(content, node.data)

	return strings.Join(content, " -> ")
}
