---
published: true
author: marshallKR7
layout: post
title:  Implementation of Merge Sorting in Go
---

## Merge Sorting

In computer science, merge sort (also commonly spelled mergesort) is an efficient, general-purpose, comparison-based sorting
algorithm. Most implementations produce a stable sort, which means that the implementation preserves the input order of equal
elements in the sorted output. Mergesort is a divide and conquer algorithm that was invented by John von Neumann in 1945.

### Analysis

#### Performance

In sorting n objects, merge sort has an average and worst-case performance of O(n log n). If the running time of merge sort
for a list of length n is T(n), then the recurrence T(n) = 2T(n/2) + n follows from the definition of the algorithm (apply 
the algorithm to two lists of half the size of the original list, and add the n steps taken to merge the resulting two lists).
The closed form follows from the master theorem for divide-and-conquer recurrences.

In the worst case, the number of comparisons merge sort makes is equal to or slightly smaller than (n ⌈lg n⌉ - 2⌈lg n⌉ + 1),
which is between (n lg n - n + 1) and (n lg n + n + O(lg n)).

#### Algorithm

Conceptually, a merge sort works as follows:
1. Divide the unsorted list into n sublists, each containing 1 element (a list of 1 element is considered sorted).
2. Repeatedly merge sublists to produce new sorted sublists until there is only 1 sublist remaining. This will be the sorted list.

## Code for implementation of Selection Sort in Go

```go
package main

import "fmt"

func Merge(a1 []int, a2 []int) []int {

  var r = make([]int, len(a1) + len(a2))
  var i = 0
  var j = 0

  for i < len(a1) && j < len(a2) {
  
    if a1[i] <= a2[j] {
      r[i+j] = a1[i]
      i++
    } else {
      r[i+j] = a2[j]
      j++
    }
    
  }

  for i < len(a1) { r[i+j] = a1[i]; i++ }
  for j < len(a2) { r[i+j] = a2[j]; j++ }

  return r
  
}

func Merge_Sort(items []int) []int {

  if len(items) < 2 {
    return items
    
  }

  var middle = len(items) / 2
  var a = Merge_Sort(items[:middle])
  var b = Merge_Sort(items[middle:])
  return Merge(a, b)
  
}

func main () {

  fmt.Print(Merge_Sort([]int{ 10, 92, 83, 43,45, 64, 744, 33, 22, 14,2344,56743 }), "\n")
  
}
```
