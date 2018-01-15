---
published: true
author: surasnayak
layout: post
title: Implementation of Binary Tree Datastructure in GO
---

## Tree Datastructure

- A tree is a collection of nodes connected by directed (or undirected) edges. 
- A tree is a nonlinear data structure, compared to arrays, linked lists, stacks and queues which are linear data structures. 
- A tree can be empty with no nodes or a tree is a structure consisting of one node called the root and zero or one or more subtrees. 

### general properties :

- One node is distinguished as a root.
- Every node (exclude a root) is connected by a directed edge from exactly one other node. 
- A direction is: parent -> children.

- Each node can have arbitrary number of children. Nodes with no children are called leaves, or external nodes. 

- Nodes with the same parent are called siblings.

### Algorithm for Implementation :

- Since each node in a tree can have an arbitrary number of children, and that number is not known in advance, the general tree can be implemented using a first child/next sibling method. 

- Each node will have TWO pointers: 

1) one to the leftmost child.
2) one to the rightmost sibling. 

## Code for implementaion of Tree data structure in GO :

```go

// This program compares a pair of trees by
// walking each in its own goroutine,
// sending their contents through a channel
// to a third goroutine that compares them.

package main

import (
	"fmt"
	"math/rand"
)

// A Tree is a binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk traverses a tree depth-first,
// sending each Value on a channel.
func Walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Walker launches Walk in a new goroutine,
// and returns a read-only channel of values.
func Walker(t *Tree) <-chan int {
	ch := make(chan int)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}

// Compare reads values from two Walkers
// that run simultaneously, and returns true
// if t1 and t2 have the same contents.
func Compare(t1, t2 *Tree) bool {
	c1, c2 := Walker(t1), Walker(t2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			break
		}
	}
	return false
}

// New returns a new, random binary tree
// holding the values 1k, 2k, ..., nk.
func New(n, k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(n) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func main() {
	t1 := New(100, 1)
	fmt.Println(Compare(t1, New(100, 1)), "Same Contents")
	fmt.Println(Compare(t1, New(99, 1)), "Differing Sizes")
	fmt.Println(Compare(t1, New(100, 2)), "Differing Values")
	fmt.Println(Compare(t1, New(101, 2)), "Dissimilar")
}

```