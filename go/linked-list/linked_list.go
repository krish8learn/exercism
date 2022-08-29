package linkedlist

import "fmt"

// Define List and Node types here.
// Note: The tests expect Node type to include an exported field with name Value to pass.

type Node struct {
	prev  *Node
	Value interface{}
	next  *Node
}

type List struct {
	head *Node
	tail *Node
}

func NewList(args ...interface{}) *List {
	// panic("Please implement the NewList function")
	l := &List{}
	previousNodeRemember := &Node{}
	for index, value := range args {
		//present node
		presentNode := &Node{Value: value}
		if index == 0 {
			//first element
			l.head = presentNode
		} else if index == len(args)-1 {
			//last element
			presentNode.prev = previousNodeRemember
			previousNodeRemember.next = presentNode
			l.tail = presentNode
		} else {
			//middle element
			presentNode.prev = previousNodeRemember
			previousNodeRemember.next = presentNode
		}

		if len(args) == 1 {
			//total element entry is 1
			l.head = presentNode
			l.tail = presentNode
		}
		previousNodeRemember = presentNode
	}

	return l
}

func (n *Node) Next() *Node {
	// panic("Please implement the Next function")
	return n.next
}

func (n *Node) Prev() *Node {
	// panic("Please implement the Prev function")
	return n.prev
}

func (l *List) Unshift(v interface{}) {
	// panic("Please implement the Unshift function")

	newNode := &Node{Value: v}
	// list is empty
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		//list is not empty
		l.head, newNode.next, l.head.prev = newNode, l.head, newNode
	}

}

func (l *List) Push(v interface{}) {
	// panic("Please implement the Push function")
	//insert data at the end of the linked list

	//new node
	newNode := &Node{prev: l.tail, Value: v, next: nil}

	//if the list is empty
	if l.tail == nil || l.head == nil {
		l.head = newNode
		l.tail = newNode
	}
	//list isn't empty
	tempNode := l.tail
	tempNode.next = newNode
	l.tail = tempNode.next
}

func (l *List) Shift() (interface{}, error) {
	// panic("Please implement the Shift function")
	//remove value from the front

	var returnValue interface{}

	//list is empty
	if l.head == nil {
		return returnValue, fmt.Errorf("list is empty")
	}

	//list is not empty
	returnValue = l.head.Value
	l.head = l.head.next

	if l.head != nil {
		l.head.prev = nil
	} else {
		//list is becoming empty
		l.tail = nil
	}
	return returnValue, nil
}

func (l *List) Pop() (interface{}, error) {
	// panic("Please implement the Pop function")

	var returnValue interface{}
	//list is empty
	if l.tail == nil {
		return returnValue, fmt.Errorf("list is empty")
	}

	//list is not empty
	returnValue = l.tail.Value
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
		//l.tail.prev = nil
	} else {
		//list becoming empty
		l.head = nil
	}

	return returnValue, nil
}

func (l *List) Reverse() {
	// panic("Please implement the Reverse function")
	l.head, l.tail = l.tail, l.head

	tempNode := l.head

	for tempNode != nil {
		tempNode.prev, tempNode.next = tempNode.next, tempNode.prev
		tempNode = tempNode.next
	}

}

func (l *List) First() *Node {
	// panic("Please implement the First function")
	return l.head
}

func (l *List) Last() *Node {
	// panic("Please implement the Last function")
	return l.tail
}
