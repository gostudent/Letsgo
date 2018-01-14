---
published: true
---

## Rank of all elements in an array

### Question:

Given an array of N integers with duplicates allowed. All elements are ranked from 1 to N if they are distinct. If there are say x repeated elements of a particular value then each element should be assigned a rank equal to the arithmetic mean of x consecutive ranks.


### Sample Input:

The first line of the input contains a single integer N size of the array.
Next line contains N-Space separated integers, denoting elements of an array.

### Output:

For each element of the array, print rank of that element .


#### Sample TestCase 1
###### Input :
- 3
- 20 30 10 

###### Output :

- 2.0 3.0 1.0

#### Sample TestCase 2
###### Input :
- 7
- 1 2 5 2 1 60 3


###### Output :

- 1.5 3.5 6.0 3.5 1.5 7.0 5.0



## Solution

- **Language Used:** go

```go
package main

import (
	"fmt"
)

func Bubblesort(arr [3]int) {

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

func Swap(arr [3]int, i, j int) {
	tmp := arr[j]
	arr[j] = arr[i]
	arr[i] = tmp
}

func main() {
	var n , i, j int;
	fmt.Scan(&n)
	var arr , brr [3]int
	for i = 0; i<n; i++ {
		fmt.Scan(&arr[i])
		brr[i] = arr[i];
	}
	Bubblesort(arr)
	for i =0; i< n; {
		c := arr[i]
		var val float64  =0;
		var cnt int = 0;
		for j = i; j <n; j++ {
			if(c == arr[j] ){
				val = val + float64(j);
				cnt++;
				i++;
			} else{
				j = n;
			}
		}
		val = val + float64(cnt);
		val = val/ float64(cnt);
		for j =0; j <n; j++{
			if ( c == brr[j] ){
				brr[j] = c;
			}
		}
	}
	for i =0; i<n; i++{
		fmt.Println(brr[i], " "); 
	}

}
```

