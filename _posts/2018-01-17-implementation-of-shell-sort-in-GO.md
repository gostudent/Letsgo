---
published: true
author: surasnayak
layout: post
title: Implementation of Shell Sort in GO
---

## Shell Sort :

- Shell sort is a highly efficient sorting algorithm and is based on insertion sort algorithm. This algorithm avoids large shifts as in case of insertion sort, if the smaller value is to the far right and has to be moved to the far left.

- This algorithm uses insertion sort on a widely spread elements, first to sort them and then sorts the less widely spaced elements. This spacing is termed as interval. 
- This interval is calculated based on Knuth's formula as Knuth's Formula

### Knuth's Formula :

h = h * 3 + 1

where −
   h is interval with initial value 1



 ### Algorithm for shell sort :

- Step 1 − Initialize the value of h
- Step 2 − Divide the list into smaller sub-list of equal interval h
- Step 3 − Sort these sub-lists using insertion sort
- Step 4 − Repeat until complete list is sorted


### Time Complexity for shell sort :


- Shellsort has prospects of running in an average time that asymptotically grows like NlogN only when using gap sequences whose number of gaps grows in proportion to the logarithm of the array size. 
- It is, however, unknown whether Shellsort can reach this asymptotic order of average-case complexity, which is optimal for comparison sorts.

## Code for implementaion of Shell Sort in GO :

```go

package main

import "fmt"
import "math/rand"
func main() {
	arrayz := Random_Arr(10)
    fmt.Println("Initial array is as follows:", arrayz)
    fmt.Println("")

    for d := int(len(arrayz)/2); d > 0; d /= 2 {
    	for j := d; j < len(arrayz); j++ {
    		for k := j; k >= d && arrayz[k-d] > arrayz[k]; k -= d {
    			arrayz[k], arrayz[k-d] = arrayz[k-d], arrayz[k]
    		}
    	}
    }

    fmt.Println("Sorted array is as follows: ", arrayz) 	
}
func Random_Arr(n int) []int {
    arrayz := make([]int, n)
    for i := 0; i <= n - 1; i++ {
        arrayz[i] = rand.Intn(n)
    }
    return arrayz
}

```