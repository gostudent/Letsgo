---
published: true
author: surasnayak
layout: post
title: Implementation of Quick Sort in GO
---

## Quick sort :

- Like Merge Sort, QuickSort is a Divide and Conquer algorithm.
- It picks an element as pivot and partitions the given array around the picked pivot. 
- There are many different versions of quickSort that pick pivot in different ways.

- Always pick first element as pivot.
- Always pick last element as pivot (implemented below)
- Pick a random element as pivot.
- Pick median as pivot.
- The key process in quickSort is partition(). 
- Target of partitions is, given an array and an element x of array as pivot, put x at its correct position in sorted array and put all smaller elements (smaller than x) before x, and put all greater elements (greater than x) after x. 
- All this should be done in linear time.

### Partition Algorithm :

- There can be many ways to do partition, following pseudo code adopts the method given in CLRS book. 
- The logic is simple, we start from the leftmost element and keep track of index of smaller (or equal to) elements as i. 
- While traversing, if we find a smaller element, we swap current element with arr[i]. 
- Otherwise we ignore current element.

### Quick Sort Algorithm :

- Using pivot algorithm recursively, we end up with smaller possible partitions. 
- Each partition is then processed for quick sort.
- We define recursive algorithm for quicksort as follows −

Step 1 − Make the right-most index value pivot
Step 2 − partition the array using pivot value
Step 3 − quicksort left partition recursively
Step 4 − quicksort right partition recursively

### Time Complexity of quick sort :

- The first two terms are for two recursive calls, the last term is for the partition process. 
- k is the number of elements which are smaller than pivot.
- The time taken by QuickSort depends upon the input array and partition strategy. 
- Following are three cases :

- Worst Case: 
- The worst case occurs when the partition process always picks greatest or smallest element as pivot. 
- If we consider above partition strategy where last element is always picked as pivot, the worst case would occur when the array is already sorted in increasing or decreasing order. 
- Following is recurrence for worst case.

T(n) = T(0) + T(n-1) + theta(n)
which is equivalent to  
 T(n) = T(n-1) + theta(n)
 
 - The solution of above recurrence is theta(n2).
 
 - Best Case:
 - The best case occurs when the partition process always picks the middle element as pivot. 
 - Following is recurrence for best case.

 T(n) = 2T(n/2) + theta(n)
 
- The solution of above recurrence is theta(nLogn).


## Code for implementaion of Quick sort in GO :

```go
package main

import "fmt"
import "math/rand"


func Quick_Sorting(arrayz []int) []int {
    
    if len(arrayz) <= 1 {
        return arrayz
    }
    median := arrayz[rand.Intn(len(arrayz))]
    low := make([]int, 0, len(arrayz))
    high := make([]int, 0, len(arrayz))
    middle := make([]int, 0, len(arrayz))
    for _, item := range arrayz {
        switch {
        case item < median:
            low = append(low, item)
        case item == median:
            middle = append(middle, item)
        case item > median:
            high = append(high, item)
        }
    }
 
    low = Quick_Sorting(low)
    high = Quick_Sorting(high)

    low = append(low, middle...)
    low = append(low, high...)
 
    return low
}

func main() {
    arrayz := RandomArr(10)

    fmt.Println("Initial array :", arrayz)
    fmt.Println("")
    fmt.Println("Sorted array : ", Quick_Sorting(arrayz))
}

func RandomArr(n int) []int {
    arrayz := make([]int, n)
    for i := 0; i <= n - 1; i++ {
        arrayz[i] = rand.Intn(n)
    }
    return arrayz
}

```
