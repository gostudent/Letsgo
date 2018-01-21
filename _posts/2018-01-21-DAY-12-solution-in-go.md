---
published: true
author: ADESH SHUKLA
layout: post
title: Solve The problem!
---

## Question :

The problem is like this: You are given an array, A of size N. You have to find the strength of all the elements of the array. And the strength of ith element Ai is the number of elements Aj smaller than Ai such that j>i.

## Input :

The first line of the input contains a single integer, N denoting the size of the array. And the second line contains N space-separated positive integers where ith integer is the ith elements of array.

## Output :

Print N space-separated integers denoting the strength of each element of the array.

## Sample TestCase :

### Input :
5
5 3 4 3 2
### Output :
4 1 2 1 0
## Solution :

```go

package main
import "fmt"

func main(){
			
			var n int
			fmt.Scan(&n)
			var arr [100007]int
			var cn [100007]int
			for j:=0; j<n; j++ {
				fmt.Scan(&arr[j])
			}
			
			for j:=0; j<n; j++ {
				cn[j]=0
				for k:=j+1; k<n; k++ {
					
					if arr[k] < arr[j] {
						cn[j]++
						}
					}
				}
			
			
			for j:=0; j<n; j++ {
				fmt.Print(cn[j])
				fmt.Print(" ")

			}
		
}
``