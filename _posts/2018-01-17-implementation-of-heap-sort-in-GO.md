---
published: true
author: surasnayak
layout: post
title: Implementation of Heap in GO
---

## Heap Sort

- Heap sort is a comparison based sorting technique based on Binary Heap data structure. 
-It is similar to selection sort where we first find the maximum element and place the maximum element at the end. 
- We repeat the same process for remaining element.

### What is Binary Heap?

- Let us first define a Complete Binary Tree. 
- A complete binary tree is a binary tree in which every level, except possibly the last, is completely filled, and all nodes are as far left as possible (Source Wikipedia)

- A Binary Heap is a Complete Binary Tree where items are stored in a special order such that value in a parent node is greater(or smaller) than the values in its two children nodes. 
- The former is called as max heap and the latter is called min heap. 
- The heap can be represented by binary tree or array.

### Heap Sort Algorithm for sorting in increasing order:

1. Build a max heap from the input data.
2. At this point, the largest item is stored at the root of the heap. Replace it with the last item of the heap followed by reducing the size of heap by 1. Finally, heapify the root of tree.
3. Repeat above steps while size of heap is greater than 1.

### How to build the heap?

- Heapify procedure can be applied to a node only if its children nodes are heapified. 
- So the heapification must be performed in the bottom up order.

### Time Complexity for Heap Sort: 

max_heapify has complexity O(logN), build_maxheap has complexity O(N) and we run max_heapify 
N−1 times in heap_sort function, therefore complexity of heap_sort function is O(NlogN).

## Code for implementaion of queue data structure in GO :

```go
package main

import "fmt"
import "math/rand"
import "time"

type MaxHeap struct {
	slice    []int
	heapSize int
}

func BuildMaxHeap(slice []int) MaxHeap {
	h := MaxHeap{slice: slice, heapSize: len(slice)}
	for i := len(slice) / 2; i >= 0; i-- {
		h.MaxHeapify(i)
	}
	return h
}

func (h MaxHeap) MaxHeapify(i int) {
	l, r := 2*i+1, 2*i+2
	max := i

	if l < h.size() && h.slice[l] > h.slice[max] {
		max = l
	}
	if r < h.size() && h.slice[r] > h.slice[max] {
		max = r
	}
	if max != i {
		h.slice[i], h.slice[max] = h.slice[max], h.slice[i]
		h.MaxHeapify(max)
	}
}

func (h MaxHeap) size() int { return h.heapSize } 

func heapSort(slice []int) []int {
	h := BuildMaxHeap(slice)
	for i := len(h.slice) - 1; i >= 1; i-- {
		h.slice[0], h.slice[i] = h.slice[i], h.slice[0]
		h.heapSize--
		h.MaxHeapify(0)
		if i == len(h.slice)-1 || i == len(h.slice)-3 || i == len(h.slice)-5 {
			element := (i - len(h.slice)) * -1
			fmt.Println("Heap after removing ", element, " elements")
			fmt.Println(h.slice)

		}
	}
	return h.slice
}

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	s := make([]int, 20)
	for index, _ := range s {
		s[index] = r1.Intn(400)
	}
	fmt.Println("Randomly generated array:")
	fmt.Println(s)
	h := BuildMaxHeap(s)
	fmt.Println("\nInital Heap: ")
	fmt.Println(h.slice, "\n")

	s = heapSort(s)
	fmt.Println("\nFinal List:")
	fmt.Println(s)
}

```
