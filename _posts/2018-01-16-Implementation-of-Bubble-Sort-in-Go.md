---
published: true
author: marshallKR7
layout: post
title:  Implementation of Bubble Sorting in Go
---

## Bubble Sorting

Bubble sort is a simple sorting algorithm that repeatedly steps through the list to be sorted, compares each pair of adjacent 
items and swaps them if they are in the wrong order. The pass through the list is repeated until no swaps are needed, which 
indicates that the list is sorted. The algorithm, which is a comparison sort, is named for the way smaller or larger elements 
"bubble" to the top of the list. Although the algorithm is simple, it is too slow and impractical for most problems even when
compared to insertion sort.It can be practical if the input is usually in sorted order but may occasionally have some out-of-order
elements nearly in position.

### Analysis

#### Performance

Bubble sort has worst-case and average complexity both О(n2), where n is the number of items being sorted.There exist many sorting
algorithms, such as merge sort with substantially better worst-case or average complexity of O(n log n).Therefore, bubble sort is 
not a practical sorting algorithm when n is large.
The only significant advantage that bubble sort has over most other implementations, even quicksort, but not insertion sort, 
is that the ability to detect that the list is sorted efficiently is built into the algorithm. When the list is already sorted
(best-case), the complexity of bubble sort is only O(n). By contrast, most other algorithms, even those with better average-case
complexity, perform their entire sorting process on the set and thus are more complex.
Bubble sort should be avoided in the case of large collections. It will not be efficient in the case of a reverse-ordered 
collection.

#### Step-by-step example

Let us take the array of numbers "5 1 4 2 8", and sort the array from lowest number to greatest number using bubble sort. 
In each step, elements written in bold are being compared. Three passes will be required.

**First Pass**
- ( **5** **1** 4 2 8 ) → ( **1** **5** 4 2 8 ), Here, algorithm compares the first two elements, and swaps since 5 > 1.
- ( 1 **5** **4** 2 8 ) → ( 1 **4** **5** 2 8 ), Swap since 5 > 4
- ( 1 4 **5** **2** 8 ) → ( 1 4 **2** **5** 8 ), Swap since 5 > 2
- ( 1 4 2 **5** **8** ) → ( 1 4 2 **5** **8** ), Now, since these elements are already in order (8 > 5), algorithm does not swap them.

**Second Pass**
- ( **1** **4** 2 5 8 ) → ( **1** **4** 2 5 8 )
- ( 1 **4** **2** 5 8 ) → ( 1 **2** **4** 5 8 ), Swap since 4 > 2
- ( 1 2 **4** **5** 8 ) → ( 1 2 **4** **5** 8 )
- ( 1 2 4 **5** **8** ) → ( 1 2 4 **5** **8** )
- Now, the array is already sorted, but the algorithm does not know if it is completed. The algorithm needs one whole pass 
 without any swap to know it is sorted.
 
 **Third Pass**
- ( **1** **2** 4 5 8 ) → ( **1** **2** 4 5 8 )
- ( 1 **2** **4** 5 8 ) → ( 1 **2** **4** 5 8 )
- ( 1 2 **4** **5** 8 ) →( 1 2 **4** **5** 8 )
- ( 1 2 4 **5** **8** ) → ( 1 2 4 **5** **8** )

## Code for implementation of Bubble Sort in Go

```go
package main

import (
	"fmt"
)

func Bubblesort(arr []int) {

	swapped := true;
	for swapped {
		swapped = false
		for i := 0; i < len(arr) - 1; i++ {
			if arr[i + 1] < arr[i] {
				Swap(arr, i, i + 1)
				swapped = true
			}
		}
	}	
}

func Swap(arr []int, i, j int) {
	tmp := arr[j]
	arr[j] = arr[i]
	arr[i] = tmp
}

func main() {

	arr := []int{199,66, 52,44, 29, 40, 45, 33, 74, 34}
	fmt.Println("Unsorted array is as follows: ", arr)
	Bubblesort(arr)
	fmt.Println("Sorted array is as follows: ", arr)
}
```
