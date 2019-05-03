package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

var size = 0
var stack = new(Node)

func Push(v int) bool {
	if stack == nil {
		stack = &Node{v, nil}
		size = 1
		return true
	}

	temp := &Node{v, nil}
	temp.Next = stack
	stack = temp
	size++
	return true
}

func Pop(t *Node) (int, bool) {
	if size == 0 {
		return 0, false
	}

	if size == 1 {
		size = 0
		stack = nil
		return t.Value, true
	}

	stack = stack.Next
	size--
	return t.Value, true
}

func traverse(t *Node) {
	if size == 0 {
		fmt.Println("Empty Stack!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {

	stack = nil
	v, b := Pop(stack)
	if b {
		fmt.Print(v, " ")
	} else {
		fmt.Println("Pop() failed!")
	}

	Push(100)
	traverse(stack)
	Push(200)
	traverse(stack)

	for i := 0; i < 10; i++ {
		Push(i)
	}

	for i := 0; i < 15; i++ {
		v, b := Pop(stack)
		if b {
			fmt.Print(v, " ")
		} else {
			break
		}
	}
	fmt.Println()
	traverse(stack)
}