package binarysearchtree

import (
	"sort"
)

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	// panic("Please implement the NewBst function")
	return &BinarySearchTree{
		data: i,
	}
}

// Insert inserts an int into the BinarySearchTree.
// Inserts happen based on the rules of a binary search tree
func (bst *BinarySearchTree) Insert(i int) {
	// panic("Please implement the Insert method")

	if i <= bst.data {
		if bst.left == nil {
			bst.left = NewBst(i)
		} else {
			bst.left.Insert(i)
		}
	} else if i > bst.data {
		if bst.right == nil {
			bst.right = NewBst(i)
		} else {
			bst.right.Insert(i)
		}
	}
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
// The values are in increasing order starting with the lowest int value.
// A BinarySearchTree that has the numbers [1,3,7,5] added will return the
// []int [1,3,5,7].
func (bst *BinarySearchTree) SortedData() []int {
	// panic("Please implement the SortedData function")

	result := make([]int, 0)
	if bst == nil {
		return result
	} else {
		leftData := bst.left.SortedData()
		if len(leftData) != 0 {
			result = append(result, leftData...)
		}
		rightData := bst.right.SortedData()
		if len(rightData) != 0 {
			result = append(result, rightData...)
		}
	}
	result = append(result, bst.data)
	sort.Ints(result)
	
	return result
}
