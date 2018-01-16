---
published: true
author: surasnayak
layout: post
title: Implementation of Stack Datastructure in GO
---

## Stack Datastructure

- Stack is an abstract data type with a bounded(predefined) capacity. 
- It is a simple data structure that allows adding and removing elements in a particular order. 
- Every time an element is added, it goes on the top of the stack, the only element that can be removed is the element that was at the top of the stack, just like a pile of objects.

### Basic features of Stack :

- Stack is an ordered list of similar data type.
- Stack is a LIFO structure. (Last in First out).
- push() function is used to insert new elements into the Stack and pop() function is used to delete an element from the stack. Both insertion and deletion are allowed at only one end of Stack called Top.
- Stack is said to be in Overflow state when it is completely full and is said to be in Underflow state if it is completely empty. 

### Applications of Stack :

- The simplest application of a stack is to reverse a word. You push a given word to stack - letter by letter - and then pop letters from the stack.

- here are other uses also like : Parsing, Expression Conversion(Infix to Postfix, Postfix to Prefix etc) and many more.

### Algorith for PUSH operation:

- Check if the stack is full or not.
- If the stack is full,then print error of overflow and exit the program.
- If the stack is not full,then increment the top and add the element .

### Algorith for POP operation:

- Check if the stack is empty or not.
- If the stack is empty, then print error of underflow and exit the program.
- If the stack is not empty, then print the element at the top and decrement the top.


## Code for implementaion of Stack data structure in GO :

```go

package main

import "fmt"

type Stack struct {
	size int
	top *Node
}

type Node struct {
	value string
	next *Node
}

func (s *Stack) Length() int {
  return s.size
}

func (s *Stack) Push(val string) {
	s.top = &Node{val, s.top}
	fmt.Printf("%v pushed to stack\n",val)
	s.size++
}

func (s *Stack) Pop() (val string) {
  if s.size > 0 {
    val, s.top = s.top.value, s.top.next
    s.size--
    return 
  }
  return
}

func main() {
	stack := &Stack{}
	
	stack.Push("10")
	stack.Push("20")
	stack.Push("30")
	
	fmt.Println()
	
	for stack.Length() > 0 {
		fmt.Printf("%s popped from stack\n", stack.Pop())
	}
}

```