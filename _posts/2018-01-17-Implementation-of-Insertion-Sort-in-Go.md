---
published: true
author: marshallKR7
layout: post
title:  Implementation of Insertion Sorting in Go
---

## Insertion Sorting

Insertion sort is a simple sorting algorithm that builds the final sorted array (or list) one item at a time. It is much less
efficient on large lists than more advanced algorithms such as quicksort, heapsort, or merge sort.

However, insertion sort provides several advantages:
- Simple implementation: Jon Bentley shows a three-line C version, and a five-line optimized version.
- Efficient for (quite) small data sets, much like other quadratic sorting algorithms
- More efficient in practice than most other simple quadratic (i.e., O(n2)) algorithms such as selection sort or bubble sort
- Adaptive, i.e., efficient for data sets that are already substantially sorted: the time complexity is O(nk) when each element in the input is no more than k places away from its sorted position
- Stable; i.e., does not change the relative order of elements with equal keys
- In-place; i.e., only requires a constant amount O(1) of additional memory space
- Online; i.e., can sort a list as it receives it

When people manually sort cards in a bridge hand, most use a method that is similar to insertion sort.

### Analysis

#### Performance
- The best case input is an array that is already sorted. In this case insertion sort has a linear running time (i.e., O(n)).
During each iteration, the first remaining element of the input is only compared with the right-most element of the sorted 
subsection of the array.
- The simplest worst case input is an array sorted in reverse order. The set of all worst case inputs consists of all arrays
where each element is the smallest or second-smallest of the elements before it. In these cases every iteration of the inner
loop will scan and shift the entire sorted subsection of the array before inserting the next element. This gives insertion sort
a quadratic running time (i.e., O(n2)).
- The average case is also quadratic, which makes insertion sort impractical for sorting large arrays. However, insertion sort
is one of the fastest algorithms for sorting very small arrays, even faster than quicksort; indeed, good quicksort implementations
use insertion sort for arrays smaller than a certain threshold, also when arising as subproblems; the exact threshold must be 
determined experimentally and depends on the machine, but is commonly around ten.

#### Example
 The following table shows the steps for sorting the sequence {3, 7, 4, 9, 5, 2, 6, 1}. In each step, the key under 
 consideration is shown in italics. The key that was moved (or left in place because it was biggest yet considered) in the
 previous step is shown in bold.
- *3* 7 4 9 5 2 6 1
- **3** *7* 4 9 5 2 6 1
- 3 **7** *4* 9 5 2 6 1
- 3 **4** 7 *9* 5 2 6 1
- 3 4 7 **9** *5* 2 6 1
- 3 4 **5** 7 9 *2* 6 1
- **2** 3 4 5 7 9 *6* 1
- 2 3 4 5 **6** 7 9 *1*
- **1** 2 3 4 5 6 7 9

## Code for implementation of Selection Sort in Go

```go
package main


import "fmt"


func main() {
	arrayz := [9]int{2,1,4,3,5,9,7,6,8}

	for i:=1; i<len(arrayz); i++{
		tem:=arrayz[i]
		j := i

		for;j>0 && arrayz[j-1]>=tem;j--{
			arrayz[j] = arrayz[j-1]
		}
		arrayz[j] = tem
	}

	for Sorted_val:= range arrayz{
		fmt.Println(Sorted_val)
	}
}
```

