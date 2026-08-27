package exercises

// Node is a node in a singly linked list. The Next *Node self-reference is only
// possible through a pointer: a struct cannot contain itself by value (that would
// require infinite size). nil Next marks the end of the list.
type Node struct {
	Value int
	Next  *Node
}

// Append adds value to the END of the list and returns the head of the list.
// If head is nil (empty list), the new node becomes the head — that is why
// Append returns the head rather than modifying it in place.
//
// TODO: implement. Walk to the last node and attach a new one; handle nil head.
func Append(head *Node, value int) *Node {
	// TODO: implement
	return nil
}

// Values walks the list from head and returns the values in order. A nil head
// yields a nil slice.
//
// TODO: implement by following Next until it is nil.
func Values(head *Node) []int {
	// TODO: implement
	return nil
}
