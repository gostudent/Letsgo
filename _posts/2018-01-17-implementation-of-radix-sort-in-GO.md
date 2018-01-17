---
published: true
author: surasnayak
layout: post
title: Implementation of Radix Sort in GO
---

## Radix Sort :

- The lower bound for Comparison based sorting algorithm (Merge Sort, Heap Sort, Quick-Sort .. etc) is Ω(nLogn), i.e., they cannot do better than nLogn.

- Counting sort is a linear time sorting algorithm that sort in O(n+k) time when elements are in range from 1 to k.

- What if the elements are in range from 1 to n2? 

We can’t use counting sort because counting sort will take O(n2) which is worse than comparison based sorting algorithms. 

- Can we sort such an array in linear time?

- Radix Sort is the answer. The idea of Radix Sort is to do digit by digit sort starting from least significant digit to most significant digit. 
- Radix sort uses counting sort as a subroutine to sort.

### Prerequisite:

QuickSort, MergeSort, HeapSort are comparison based sorting algorithms.
CountSort is not comparison based algorithm. 
- It has the complexity ofO(n+k), where k is the maximum element of the input array.
So, if O(n) ,CountSort becomes linear sorting, which is better than comparison based sorting algorithms that have O(nlogn) time complexity. 
- he idea is to extend the CountSort algorithm to get a better time complexity when k goes O(n2). Here comes the idea of Radix Sort.

### Algorithm for Radix sort :

- For each digit i where i varies from the least significant digit to the most significant digit of a number

- Sort input array using countsort algorithm according to ith digit.

- We used count sort because it is a stable sort.


### Time Complexity for shell sort :

- Let there be d digits in input integers. Radix Sort takes O(d*(n+b)) time where b is the base for representing numbers, for example, for decimal system, b is 10. 
- What is the value of d? If k is the maximum possible value, then d would be O(logb(k)). So overall time complexity is O((n+b) * logb(k)). Which looks more than the time complexity of comparison based sorting algorithms for a large k. Let us first limit k. Let k <= nc where c is a constant. In that case, the complexity becomes O(nLogb(n)). But it still doesn’t beat comparison based sorting algorithms.

- What if we make value of b larger?. What should be the value of b to make the time complexity linear? If we set b as n, we get the time complexity as O(n). 
- In other words, we can sort an array of integers with range from 1 to nc if the numbers are represented in base n (or every digit takes log2(n) bits).

## Code for implementaion of Radix Sort in GO :

```go

package main

import (
	"fmt"
	"strconv"
)

// Finds the largest number in an array
func findLargestNum(array []int) int {
	largestNum := 0

	for i := 0; i < len(array); i++ {
		if array[i] > largestNum {
			largestNum = array[i]
		}
	}
	return largestNum
}

// Radix Sort
func radixSort(array []int) []int {

  fmt.Println("Running Radix Sort on Unsorted List\n")

  // Base 10 is used
	largestNum := findLargestNum(array)
	size := len(array)
	significantDigit := 1
	semiSorted := make([]int, size, size)

	// Loop until we reach the largest significant digit
	for largestNum / significantDigit > 0 {
  
    fmt.Println("\tSorting: " + strconv.Itoa(significantDigit) + "'s place", array)

		bucket := [10]int{0}

		// Counts the number of "keys" or digits that will go into each bucket
		for i := 0; i < size; i++ {
			bucket[(array[i] / significantDigit) % 10]++
		}

		// Add the count of the previous buckets
    // Acquires the indexes after the end of each bucket location in the array
		// Works similar to the count sort algorithm
		for i := 1; i < 10; i++ {
			bucket[i] += bucket[i - 1] 
		}

		// Use the bucket to fill a "semiSorted" array
		for i := size - 1; i >= 0; i-- {
			bucket[(array[i] / significantDigit) % 10]--
			semiSorted[bucket[(array[i] / significantDigit) % 10]] = array[i]
		}

    // Replace the current array with the semisorted array
		for i := 0; i < size; i++ {
			array[i] = semiSorted[i]
		}
    
    fmt.Println("\n\tBucket: ", bucket)
  
    // Move to next significant digit
		significantDigit *= 10 
	}

	return array
}

func main() {

  fmt.Println("\n\nRunning Radix Sort Example in Go!")
  fmt.Println("----------------------------------\n")

	unsortedList :=[]int{10, 2, 303, 4021, 293, 1, 0, 429, 480, 92, 2999, 14}
	fmt.Println("Unsorted List:", unsortedList, "\n")

	sortedList := radixSort(unsortedList)

	fmt.Println("\nSorted List:", sortedList, "\n")
}

```
