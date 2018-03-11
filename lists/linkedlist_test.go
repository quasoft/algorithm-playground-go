package goalgorithms

import (
	"testing"
)

func TestNewNode(t *testing.T) {
	node := NewNode(42)
	if node == nil {
		t.Errorf("NewNode(42) = nil, expected a struct")
	} else if node.data != 42 {
		t.Errorf("NewNode(42).data = %v, want %v", node.data, 42)
	}
}

func TestNode_String(t *testing.T) {
	tests := []struct {
		name string
		n    Node
		want string
	}{
		{"Empty node", Node{}, "0"},
		{"Some value", Node{data: 42}, "42"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.n.String(); got != tt.want {
				t.Errorf("Node.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLinkedList(t *testing.T) {
	list := NewLinkedList()
	if list == nil {
		t.Errorf("NewLinkedList() = nil, expected a struct")
	}
}

func TestLinkedList_String(t *testing.T) {
	tests := []struct {
		name string
		l    *LinkedList
		want string
	}{
		{"Empty list", &LinkedList{}, "[]"},
		{"One value", NewLinkedListFromArray(42), "[42]"},
		{"Two values", NewLinkedListFromArray(42, 24), "[42, 24]"},
		{"Four values", NewLinkedListFromArray(42, 24, 256, 1024), "[42, 24, 256, 1024]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.String(); got != tt.want {
				t.Errorf("LinkedList.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Add(t *testing.T) {
	list := NewLinkedList()

	first := NewNode(42)
	second := NewNode(24)
	third := NewNode(256)

	list.head = first
	first.next = second
	second.next = third

	wantValue := 1024

	forth := list.Add(wantValue)
	if third.next != forth {
		t.Errorf("third.next was not changed to point to the newly added forth element (1024)")
	}

	if forth.data != wantValue {
		t.Errorf("list.Add(%v).data = %v, want %v", wantValue, forth.data, wantValue)
	}

	if list.head != first {
		t.Errorf("list.Add(%v) broke the head pointer of the list", wantValue)
	}

	if first.next != second {
		t.Errorf("list.Add(%v) broke the link between first and second node", wantValue)
	}

	if second.next != third {
		t.Errorf("list.Add(%v) broke the link between second and third node", wantValue)
	}

	if first.data != 42 {
		t.Errorf("list.Add(%v) somehow changed the value of the first element", wantValue)
	}

	if second.data != 24 {
		t.Errorf("list.Add(%v) somehow changed the value of the second element", wantValue)
	}

	if third.data != 256 {
		t.Errorf("list.Add(%v) somehow changed the value of the third element", wantValue)
	}
}

func TestLinkedList_Insert_BeforeFirst(t *testing.T) {
	list := NewLinkedList()

	first := NewNode(42)
	second := NewNode(24)
	third := NewNode(256)

	list.head = first
	first.next = second
	second.next = third

	value := 1024

	beforeFirst, error := list.Insert(list.head, value)
	if error != nil {
		t.Errorf("%v.Insert(%v, %v) failed with error: %v", list, list.head, value, error)
	}

	if list.head != beforeFirst {
		t.Fatalf("%v.Insert(%v, %v) should have inserted the node at the head of the list", list, list.head, value)
	}

	if beforeFirst.next != first {
		t.Errorf("%v.Insert(%v, %v) should have linked new to first node", list, list.head, value)
	}

	if first.next != second {
		t.Errorf("%v.Insert(%v, %v) should not have touched the link between first and second node", list, list.head, value)
	}
}

func TestLinkedList_Insert_BeforeThird(t *testing.T) {
	list := NewLinkedList()

	first := NewNode(42)
	second := NewNode(24)
	third := NewNode(256)

	list.head = first
	first.next = second
	second.next = third

	value := 1024

	beforeThird, error := list.Insert(third, value)
	if error != nil {
		t.Errorf("%v.Insert(%v, %v) failed with error: %v", list, third, value, error)
	}

	if second.next != beforeThird {
		t.Fatalf("%v.Insert(%v, %v) should have inserted the value before the third node", list, third, value)
	}

	if beforeThird.next != third {
		t.Errorf("%v.Insert(%v, %v) should have linked new to third node", list, third, value)
	}

	if list.head != first {
		t.Errorf("%v.Insert(%v, %v) should not have touched head of linked list", list, third, value)
	}

	if first.next != second {
		t.Errorf("%v.Insert(%v, %v) should not have touched the link between first and second node", list, second, value)
	}
}

func TestLinkedList_Insert_InEmptyList(t *testing.T) {
	list := NewLinkedList()
	node, error := list.Insert(nil, 314)
	if error == nil {
		t.Errorf("%v.Insert(%v, %v) should have refused to insert an element in an empty list", list, nil, 314)
	}

	if node != nil {
		t.Errorf("%v.Insert(%v, %v) = %v, want nil, as element was not inserted", list, nil, 314, node)
	}
}

func TestLinkedList_Insert_BeforeNodeThatIsNotInList(t *testing.T) {
	list := NewLinkedList()
	first := NewNode(42)
	list.head = first

	notInList := &Node{}
	node, error := list.Insert(notInList, 314)
	if error == nil {
		t.Errorf("%v.Insert(%v, %v) should have refused to insert before a node that is not in the list", list, notInList, 314)
	}

	if node != nil {
		t.Errorf("%v.Insert(%v, %v) = %v, want nil, as element was not inserted ", list, notInList, 314, node)
	}
}
