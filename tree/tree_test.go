package goalgorithms

import (
	"testing"
)

func TestNode_String(t *testing.T) {
	tests := []struct {
		name string
		t    *Node
		want string
	}{
		{"Root only", &Node{IntValue(7), nil, nil}, "(7)"},
		{"Two levels", &Node{IntValue(7), &Node{IntValue(3), nil, nil}, &Node{IntValue(17), nil, nil}}, "((3) 7 (17))"},
		{"Three levels", &Node{IntValue(7), &Node{IntValue(3), &Node{IntValue(2), nil, nil}, nil}, &Node{IntValue(17), nil, nil}}, "(((2) 3) 7 (17))"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.t.String(); got != tt.want {
				t.Errorf("Node.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_Insert(t *testing.T) {
	// Start with tree:
	//               7
	//       3               17
	//    2     6	    12         22
	//                           20  23
	//                         19
	tree := &Node{}
	tree.Value = IntValue(7)
	tree.Left = &Node{IntValue(3), &Node{IntValue(2), nil, nil}, &Node{IntValue(6), nil, nil}}
	tree.Right = &Node{IntValue(17), &Node{IntValue(12), nil, nil}, &Node{IntValue(22), nil, nil}}
	tree.Right.Right.Left = &Node{IntValue(20), &Node{IntValue(19), nil, nil}, nil}
	tree.Right.Right.Right = &Node{IntValue(23), nil, nil}

	value := 5
	tree.Insert(IntValue(value))
	want := "(((2) 3 ((5) 6)) 7 ((12) 17 (((19) 20) 22 (23))))"
	got := tree.String()
	if got != want {
		t.Errorf("tree.Insert(%v) = %v, want %v", value, got, want)
	}

	value = 4
	tree.Insert(IntValue(value))
	want = "(((2) 3 (((4) 5) 6)) 7 ((12) 17 (((19) 20) 22 (23))))"
	got = tree.String()
	if got != want {
		t.Errorf("tree.Insert(%v) = %v, want %v", value, got, want)
	}

	value = 18
	tree.Insert(IntValue(value))
	want = "(((2) 3 (((4) 5) 6)) 7 ((12) 17 ((((18) 19) 20) 22 (23))))"
	got = tree.String()
	if got != want {
		t.Errorf("tree.Insert(%v) = %v, want %v", value, got, want)
	}
}
