---
published: true
author: surasnayak
layout: post
title: Implementation of Queue Datastructure in GO
---

## Queue Datastructure

- Queue is an abstract data type or a linear data structure, in which the first element is inserted from one end called REAR(also called tail), and the deletion of existing element takes place from the other end called as FRONT(also called head). - This makes queue as FIFO(First in First Out) data structure, which means that element inserted first will also be removed first.

### Basic features of queue : 

- Like Stack, Queue is also an ordered list of elements of similar data types.
- Queue is a FIFO( First in First Out ) structure.
- Once a new element is inserted into the Queue, all the elements inserted before the new element in the queue must be removed, to remove the new element.

### Applications of queue :

- Serving requests on a single shared resource, like a printer, CPU task scheduling etc.
- In real life scenario, Call Center phone systems uses Queues to hold people calling them in an order, until a service representative is free.
- Handling of interrupts in real-time systems. The interrupts are handled in the same order as they arrive i.e First come first served.

### Algorith for ENQUEUE operation:

- Check if the queue is full or not.
- If the queue is full, then print overflow error and exit the program.
- If the queue is not full, then increment the tail and add the element.

### Algorith for DEQUEUE operation:

- Check if the queue is empty or not.
- If the queue is empty, then print underflow error and exit the program.
- If the queue is not empty, then print the element at the head and increment the head.

## Code for implementaion of queue data structure in GO :

``` go

package main

import "fmt"

type Queue struct {
	size int
	front *Node
}

type Node struct {
	value string
	next *Node
}

func (q *Queue) Length() int {
  return q.size
}

func (q *Queue) Enqueue(val string) {
	
	if(q.front == nil){
		q.front = &Node{val, nil}
	}else{
		rear := q.front
		for rear.next != nil {
		  rear = rear.next
		}
		rear.next = &Node{val, nil}
	}
	fmt.Printf("%v enqueued to queue\n",val)
	q.size++
}

func (q *Queue) Dequeue() (val string) {
  if q.size > 0 {
    val, q.front = q.front.value, q.front.next
    q.size--
    return 
  }
  return
}

func main() {
	queue := &Queue{}
	
	queue.Enqueue("10")
	queue.Enqueue("20")
	queue.Enqueue("30")
	
	fmt.Println()
	
	for queue.Length() > 0 {
		fmt.Printf("%s popped from queue\n", queue.Dequeue())
	}
}

```
