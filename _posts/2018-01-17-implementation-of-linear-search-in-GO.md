---
published: true
author: surasnayak
layout: post
title: Implementation of Linear Search in GO
---

## Linear Search

- Linear search is a very simple search algorithm. 
- In this type of search, a sequential search is made over all items one by one. 
- Every item is checked and if a match is found then that particular item is returned, otherwise the search continues till the end of the data collection.

### A simple approach is to do linear search :

- Start from the leftmost element of arr[] and one by one compare x with each element of arr[]
- If x matches with an element, return the index.
- If x doesn’t match with any of elements, return -1.

### Algorithm for Linear Search:

Linear Search ( Array A, Value x)

Step 1: Set i to 1
Step 2: if i > n then go to step 7
Step 3: if A[i] = x then go to step 6
Step 4: Set i to i + 1
Step 5: Go to Step 2
Step 6: Print Element x Found at index i and go to step 8
Step 7: Print element not found
Step 8: Exit

### Time complexity of Linear Search :

- The time complexity of above algorithm is O(n).


## Code for implementaion of Linear Search in GO :

```go

package main

import "fmt"

func linearSearch(array []int, query int) int {
	for i, item := range array {
		if item == query {
			return i
		}
	}
	return -1
}

func main() {
	
	array := []int{44, 23, 424, 633, 82, 102, 124, 148, 186, 168}
	index := linearSearch(array, 102)
	if index == -1 {
		fmt.Println("Number not found")
	} else {
		fmt.Println("Index: ", index)
		fmt.Println("array[", index, "] = ", array[index])
	}
}

```
