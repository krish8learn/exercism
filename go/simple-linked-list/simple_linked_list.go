package linkedlist

import "fmt"

// Define the List and Element types here.
type List struct {
	head   *Node
	length int
}

type Node struct {
	data    int
	pointer *Node
}

func New(dataArray []int) *List {
	// panic("Please implement the New function")
	starter := &List{
		head:   nil,
		length: 0,
	}
	temp := starter.head
	for index, value := range dataArray {
		if index == 0 {
			temp = &Node{
				data:    value,
				pointer: nil,
			}
			starter.length++
			starter.head = temp
		} else {
			for temp.pointer != nil {
				temp = temp.pointer
			}
			temp.pointer = &Node{
				data:    value,
				pointer: nil,
			}
			starter.length++
		}

	}
	return starter
}

func (l *List) Size() int {
	// panic("Please implement the Size function")
	return l.length
}

func (l *List) Push(element int) {
	// panic("Please implement the Push function")
	//LIFO
	temp := l.head
	if l.head == nil {
		l.head = &Node{
			data:    element,
			pointer: nil,
		}
		l.length++
	} else if temp != nil {
		for temp.pointer != nil {
			temp = temp.pointer
		}
		temp.pointer = &Node{
			data:    element,
			pointer: nil,
		}
		l.length++
	}

}

func (l *List) Pop() (int, error) {
	// panic("Please implement the Pop function")
	//LIFO , last element will be removed
	temp := l.head

	lastElement := 0
	//check whether it's empty or not
	if l.head == nil {
		//linked list is empty
		return 0, fmt.Errorf("linked list is empty")
	}else if l.length == 1{
		l.head = nil
		l.length--
	}else if temp != nil {
		//linked list is not empty
		length := l.length
		for index:= 0; index < l.length; index++ {
			if index == length-2{
				lastElement = temp.pointer.data
				temp.pointer = nil
				l.length--
			}
			temp = temp.pointer
		}
	}

	return lastElement, nil
}

func (l *List) Array() []int {
	// panic("Please implement the Array function")

	var array []int
	//empty array
	if l.length == 0 {
		return []int{}
	}

	temp := l.head
	for temp != nil {
		array = append(array, temp.data)
		temp = temp.pointer
	}

	return array
}

func (l *List) Reverse() *List {
	// panic("Please implement the Reverse function")

	//for empty linked list
	if l.head == nil {
		return l
	}

	//create an array of lists
	arrayFromLinkedList :=  l.Array()
	
	//rever the array
	reverseIntArray(arrayFromLinkedList)
	
	//create a new list with reversed array
	return New(arrayFromLinkedList)
}

func reverseIntArray(input []int) {
	inputLen := len(input)
	inputMid := inputLen / 2

	for i := 0; i < inputMid; i++ {
		j := inputLen - i - 1

		input[i], input[j] = input[j], input[i]
	}

}
