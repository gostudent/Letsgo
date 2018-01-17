---
published: true
author: surasnayak
layout: post
title: Implementation of Binary Search in GO
---

## Binary Search

- Binary search is a fast search algorithm with run-time complexity of Ο(log n). 
- This search algorithm works on the principle of divide and conquer. 
- For this algorithm to work properly, the data collection should be in the sorted form.

- Binary search looks for a particular item by comparing the middle most item of the collection. 
- If a match occurs, then the index of item is returned. 
- If the middle item is greater than the item, then the item is searched in the sub-array to the left of the middle item. 
- Otherwise, the item is searched for in the sub-array to the right of the middle item. 
- This process continues on the sub-array as well until the size of the subarray reduces to zero.

### Algorithm for Binary Search :

- Given an array A of n elements with values or records A0, A1, ..., An−1, sorted such that A0 ≤ A1 ≤ ... ≤ An−1, and target value T, the following subroutine uses binary search to find the index of T in A.

- Set L to 0 and R to n − 1.
- If L > R, the search terminates as unsuccessful.
- Set m (the position of the middle element) to the floor (the largest previous integer) of (L + R) / 2.
- If Am < T, set L to m + 1 and go to step 2.
- If Am > T, set R to m − 1 and go to step 2.
- Now Am = T, the search is done; return m.

- This iterative procedure keeps track of the search boundaries with the two variables. 
- Some implementations may check whether the middle element is equal to the target at the end of the procedure. 
- This results in a faster comparison loop, but requires one more iteration on average.

## Code for implementaion of Binary Search in GO :

```go

package main

func Binary_Search(array []int, target int, lowIndex int, highIndex int) int {
	if highIndex < lowIndex {
		return -1
	}
	mid := int((lowIndex + highIndex) / 2)
	if array[mid] > target {
		return Binary_Search(array, target, lowIndex, mid)
	} else if array[mid] < target {
		return Binary_Search(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}

func iterBinarySearch(array []int, target int, lowIndex int, highIndex int) int {
	startIndex := lowIndex
	endIndex := highIndex
	var mid int
	for startIndex < endIndex {
		mid = int((startIndex + endIndex) / 2)
		if array[mid] > target {
			endIndex = mid
		} else if array[mid] < target {
			startIndex = mid
		} else {
			return mid
		}
	}
	return -1
}

```
