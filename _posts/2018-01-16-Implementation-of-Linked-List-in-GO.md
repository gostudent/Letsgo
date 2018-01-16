---
published: true
author: surasnayak
layout: post
title: Implementation of Linked List Datastructure in GO
---

## Linked List Datastructure

- Like arrays, Linked List is a linear data structure. Unlike arrays, linked list elements are not stored at contiguous location; the elements are linked using pointers.


### Why Linked List Datastructure ?

- Arrays can be used to store linear data of similar types, but arrays have following limitations.
- 1) The size of the arrays is fixed: So we must know the upper limit on the number of elements in advance. Also, generally, the allocated memory is equal to the upper limit irrespective of the usage.
- 2) Inserting a new element in an array of elements is expensive, because room has to be created for the new elements and to create room existing elements have to shifted.

- For example, in a system if we maintain a sorted list of IDs in an array id[].

id[] = [1000, 1010, 1050, 2000, 2040].

- And if we want to insert a new ID 1005, then to maintain the sorted order, we have to move all the elements after 1000 (excluding 1000).
- Deletion is also expensive with arrays until unless some special techniques are used. For example, to delete 1010 in id[], everything after 1010 has to be moved.

### Advantages over array :

- 1) Dynamic size
- 2) Ease of insertion/deletion

### Drawbacks :

- 1) Random access is not allowed. We have to access elements sequentially starting from the first node. So we cannot do binary search with linked lists.
- 2) Extra memory space for a pointer is required with each element of the list.

### Insertion in Linked List :


 - A node can be added in three ways :

1) At the front of the linked list
2) After a given node.
3) At the end of the linked list.

### Deletion from Linked List :

- To delete a node from linked list, we need to do following steps.

1) Find previous node of the node to be deleted.
2) Changed next of previous node.
3) Free memory for the node to be deleted.

## ## Code for implementaion of Linked List data structure in GO :

```go

package main

import "fmt"

// value is the value of node; next is the pointer to next node 
type node struct {
	value int
	next *node
}

// first node, called head. It points from first node to last node 
var head *node = nil

func (l *node) pushFront(val int) *node {
	// if there's no nodes, head points to l (first node) 
	if head == nil {
		l.value = val
		l.next = nil
		head = l
		return l
	} else {
		// creating a new node equals to head 
		nnode := new(node)
		nnode = head
		// creating a second node with new value and `next -> nnode`
		//  is this way, nnode2 is before nnode
		nnode2 := &node {
			value: val,
			next: nnode,
		}
		// now head is equals nnode2 
		head = nnode2
		return head
	}
}

func (l *node) pushBack(val int) *node {
	// if there's no nodes, head points to l (first node) 
	if head == nil {
		l.value = val
		l.next = nil
		head = l
		return l
	} else {
		// read all list to last node 
		for l.next != nil {
			l = l.next
		}
		// allocate a new portion of memory 
		l.next = new(node)
		l.next.value = val
		l.next.next = nil
		return l
	}
}

func (l *node) popFront() *node {
	if head == nil {
		return head
	}
	// creating a new node equals to first node pointed by head 
	cpnode := new(node)
	cpnode = head.next
	
	// now head is equals cpnode (second node) 
	head = cpnode

	return head
}

func (l *node) popBack() *node {
	if head == nil {
		return head
	}
	// creating a new node equals to head 
	cpnode := new(node)
	cpnode = head
	
	// read list to the penultimate node 
	for cpnode.next.next != nil {
		cpnode = cpnode.next
	}
	// the penultimate node points to null. In this way the last node is deleted 
	cpnode.next = nil
	return head
}

func main() {
	lista := new(node)
	lista.pushBack(99).pushBack(65).pushBack(5) /* lista: 99 65 5 */
	lista.pushBack(44) /* lista: 99 65 5 44 */
	lista.pushFront(873) /* lista: 873 99 65 5 44  */
	lista.popFront() /* lista: 99 65 5 44 */
	lista.popBack() /* lista: 99 65 5 */
	
	// read the list until head is not nil 
	for head != nil {
		fmt.Printf("%d ",head.value)
		head = head.next /*head points to next node */
	}
}

```
