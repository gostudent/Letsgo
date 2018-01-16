---
published: true
author: marshallKR7
layout: post
title:  Implementation of Selection Sorting in Go
---

## Selection Sorting

In computer science, selection sort is a sorting algorithm, specifically an in-place comparison sort. It has O(n2) time 
complexity, making it inefficient on large lists, and generally performs worse than the similar insertion sort. Selection 
sort is noted for its simplicity, and it has performance advantages over more complicated algorithms in certain situations,
particularly where auxiliary memory is limited.
The algorithm divides the input list into two parts: the sublist of items already sorted, which is built up from left to right
at the front (left) of the list, and the sublist of items remaining to be sorted that occupy the rest of the list. Initially,
the sorted sublist is empty and the unsorted sublist is the entire input list. The algorithm proceeds by finding the smallest
(or largest, depending on sorting order) element in the unsorted sublist, exchanging (swapping) it with the leftmost unsorted
element (putting it in sorted order), and moving the sublist boundaries one element to the right.

### Analysis

#### Performance

Selection sort is not difficult to analyze compared to other sorting algorithms since none of the loops depend on the data in
the array. Selecting the minimum requires scanning n elements (taking n − 1 comparisons) and then swapping it into the first 
position. Finding the next lowest element requires scanning the remaining n − 1 elements and so on. 

#### Example

- Sorted sublist == ( )
- Unsorted sublist == (11, 25, 12, 22, 64)
- Least element in unsorted list == 11

- Sorted sublist ==  (11)
- Unsorted sublist == (25, 12, 22, 64)
- Least element in unsorted list == 12

- Sorted sublist == (11, 12)
- Unsorted sublist == (25, 22, 64)
- Least element in unsorted list == 22

- Sorted sublist == (11, 12, 22)
- Unsorted sublist == (25, 64)
- Least element in unsorted list == 25

- Sorted sublist == (11, 12, 22, 25)
- Unsorted sublist == (64)
- Least element in unsorted list == 64

- Sorted sublist == (11, 12, 22, 25, 64)
- Unsorted sublist == ( )

## Code for implementation of Selection Sort in Go

```go
package main

import "fmt"

func main() {
	arrayz := [10]int{1, 2, 31, 43, 12,23,44,443,56,54}
    fmt.Println("Initial array is as follows:", arrayz)
    fmt.Println("")

    var mini int = 0
    var tmp int = 0

    for i := 0; i < len(arrayz); i++ {
        mini = i
        for j := i + 1; j < len(arrayz); j++ {
            if arrayz[j] < arrayz[mini] {
                mini = j
            }
        }
		
        tmp = arrayz[i]
        arrayz[i] = arrayz[mini]
        arrayz[mini] = tmp
    }

    fmt.Println("Sorted array is as follows:    ", arrayz)
}
```


