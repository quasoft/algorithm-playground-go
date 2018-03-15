package goalgorithms

import "fmt"

type ValueInterface interface {
	Less(value interface{}) bool
}

// Binary tree data structure
type Node struct {
	Value ValueInterface
	Left  *Node
	Right *Node
}

type IntValue int

func (i IntValue) Less(value interface{}) bool {
	val := value.(IntValue)
	return int(i) < int(val)
}

// String method from https://github.com/golang/tour/blob/master/tree/tree.go
func (n *Node) String() string {
	if n == nil {
		return "()"
	}
	s := ""
	if n.Left != nil {
		s += n.Left.String() + " "
	}
	s += fmt.Sprint(n.Value)
	if n.Right != nil {
		s += " " + n.Right.String()
	}
	return "(" + s + ")"
}

// Insert method from https://github.com/golang/tour/blob/master/tree/tree.go
func (n *Node) Insert(value ValueInterface) *Node {
	if n == nil {
		return &Node{value, nil, nil}
	}

	if value.Less(n.Value) {
		n.Left = n.Left.Insert(value)
	} else {
		n.Right = n.Right.Insert(value)
	}
	return n
}
