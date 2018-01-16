---
published: true
author: surasnayak
layout: post
title: Implementation of Bucket Sort in GO
---

## Bucket Sort :

- Bucket sort is a comparison sort algorithm that operates on elements by dividing them into different buckets and then sorting these buckets individually. 
- Each bucket is sorted individually using a separate sorting algorithm or by applying the bucket sort algorithm recursively. 
- Bucket sort is mainly useful when the input is uniformly distributed over a range.

### Application of Bucket Sort :

Assume one has the following problem in front of them:

- One has been given a large array of floating point integers lying uniformly between the lower and upper bound. 
- This array now needs to be sorted. A simple way to solve this problem would be to use another sorting algorithm such as Merge sort, Heap Sort or Quick Sort. 
- However, these algorithms guarantee a best case time complexity of 
O(NlogN). However, using bucket sort, the above task can be completed in O(N) time.

### Advantages of Bucket Sort :


- The advantage of bucket sort is that once the elements are distributed into buckets, each bucket can be processed independently of the others. 
- This means that you often need to sort much smaller arrays as a follow-up step than the original array. 
- It also means that you can sort all of the buckets in parallel with one another.

### Disadvantages of Bucket Sort :

- The disadvantage is that if you get a bad distribution into the buckets, you may end up doing a huge amount of extra work for no benefit or a minimal benefit. 
- As a result, bucket sort works best when the data are more or less uniformly distributed or where there is an intelligent way to choose the buckets given a quick set of heuristics based on the input array. 

### Algorithm of Bucket Sort :

bucketSort(arr[], n)
1) Create n empty buckets (Or lists).
2) Do following for every array element arr[i].
.......a) Insert arr[i] into bucket[n*array[i]]
3) Sort individual buckets using insertion sort.
4) Concatenate all sorted buckets.


### Time Complexity of Bucket Sort :

- If we assume that insertion in a bucket takes O(1) time then steps 1 and 2 of the above algorithm clearly take O(n) time. 
- The O(1) is easily possible if we use a linked list to represent a bucket . 
- Step 4 also takes O(n) time as there will be n items in all buckets.
- The main step to analyze is step 3. This step also takes O(n) time on average if all numbers are uniformly distributed

## Code for implementaion of Bucket Sort in GO :

```go

package main

import (
	"fmt"
	"os"
	"strconv"
)


func insertionSort(array []float64) {
	for i := 0; i < len(array); i++ {
		temp := array[i]
		j := i - 1
		for ; j >= 0 && array[j] > temp; j-- {
			array[j+1] = array[j]
		}
		array[j+1] = temp
	}
}

func bucketSort(array []float64, bucketSize int) []float64 {
	var max, min float64
	for _, n := range array {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	nBuckets := int(max-min)/bucketSize + 1
	buckets := make([][]float64, nBuckets)
	for i := 0; i < nBuckets; i++ {
		buckets[i] = make([]float64, 0)
	}

	for _, n := range array {
		idx := int(n-min) / bucketSize
		buckets[idx] = append(buckets[idx], n)
	}

	sorted := make([]float64, 0)
	for _, bucket := range buckets {
		if len(bucket) > 0 {
			insertionSort(bucket)
			sorted = append(sorted, bucket...)
		}
	}

	return sorted
}

func main() {

	array :=[]float64{4, 1, 1, 3, 2, 2, 7, -6, 0}
	for _, arg := range os.Args[1:] {
		if n, err := strconv.ParseFloat(arg, 64); err == nil {
			array = append(array, n)
		}
	}
	fmt.Printf("Before Bucket Sort %v\n", array)
	array = bucketSort(array, 5)
	fmt.Printf("After Bucket Sort %v\n", array)
}

```


