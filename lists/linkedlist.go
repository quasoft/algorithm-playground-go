package goalgorithms

import (
	"fmt"
	"strconv"
)

// Node represents an element in a linked list with value (data) and pointer (next) to
// the next element in the list.
type Node struct {
	data int
	next *Node
}

// NewNode creates a new node object with the given value.
func NewNode(data int) *Node {
	node := Node{data: data}
	return &node
}

func (n Node) String() string {
	return strconv.Itoa(n.data)
}

// LinkedList represents a linked list of Node elements.
type LinkedList struct {
	head *Node
}

// NewLinkedList creates a new empty linked list
func NewLinkedList() *LinkedList {
	list := LinkedList{}
	return &list
}

// NewLinkedListFromArray createaa new linked list and initilizes it with
// the values from an integer array.
func NewLinkedListFromArray(values ...int) *LinkedList {
	list := LinkedList{}
	for _, v := range values {
		list.Add(v)
	}
	return &list
}

func (l LinkedList) String() string {
	s := ""
	node := l.head
	for node != nil {
		if s != "" {
			s += ", "
		}
		s += node.String()
		node = node.next
	}
	return "[" + s + "]"
}

// Add appends an element at the end of the linked list.
// Returns the newly added node.
func (l *LinkedList) Add(data int) *Node {
	newNode := NewNode(data)

	node := l.head
	if node == nil {
		l.head = newNode
		return newNode
	}

	if node != nil {
		for node.next != nil {
			node = node.next
		}
	}
	node.next = newNode
	return newNode
}

// Insert inserts an element to the linked list, before the specified node.
// Returns the newly created node or an error if before was not found.
func (l *LinkedList) Insert(before *Node, data int) (*Node, error) {
	newNode := NewNode(data)

	node := l.head
	if node == nil {
		return nil, fmt.Errorf("cannot insert before %v node, as linked list is empty", before)
	}

	if l.head == before {
		l.head = newNode
	} else {
		for node.next != before && node.next != nil {
			node = node.next
		}

		if node.next != before {
			return nil, fmt.Errorf("could find node %v in linked list", before)
		}
		node.next = newNode
	}

	newNode.next = before
	return newNode, nil
}
