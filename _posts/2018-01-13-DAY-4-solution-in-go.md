---
published: true
---

## Beautiful SubSequence

### Question:

Nowadays Babul is solving problems on sub-sequence. He is struck with a problem in which he has to find the longest sub-sequence in an array A of size N such that for all (i,j) where i!=j either A[i] divides A[j] or vice versa. If no such sub-sequence exists then print -1. Help him to accomplish this task.

### Constraints :

- 1<=T<=100
- 2<=N<=1000
- 1<=A[i]<=1000

### Sample Input:

The First line contains T no. of test cases.
Each Test case is of two lines.
The First line contains N size of the array.
Next line contains N-Space separated integers, denoting elements of an array.

### Sample Output:

For each T print the size of the longest sub-sequence satisfying the above criteria.

- Sample TestCase
- Input :
- 2
- 5
- 5 3 1 4 7
- 6
- 2 4 6 1 3 11

### Output :

- 2
- 3

### Explanation:

- First Test Case :
- Longest Sub Sequence are {5,1} , {4,1}, {3,1} etc. so size is 2.

- Second Test Case :
- Longest Sub Sequence are {1, 2, 6}, {1, 3, 6} so size is 3.

## Solution

- **Language Used:** go

package main
import "fmt"

func main(){
	
	var t int
	fmt.Scan(&t)
	for z:=0; z<t; z++ {
		var n int
		fmt.Scan(&n)
		var a [100009]int
		
		for i:=0; i<n; i++ {
			fmt.Scan(&a[i])
		}
		for i:=0; i<n; i++ {
			for j:=i+1; j<n; j++ {
				if(a[i] > a[j]) {
					var tmp int
					tmp=a[i]
					a[i]=a[j]
					a[j]=tmp
				}
			}
		}
		
		var b [100009]int
		
		for i:=0; i<n; i++ {
			b[i]=1
		}
		
		var max_max=0
		
		for i:=1; i<n; i++ {
			var tmp int
			tmp=a[i]
			for j:=0; j<i; j++ {
				if(a[j]%tmp==0 || tmp%a[j]==0) {
					if(b[i] > (b[j]+1)) {
						b[i]=b[i]
					} else {
						b[i]=b[j]+1
					}
					
					if(max_max > b[i]) {
						max_max=max_max
					} else {
						max_max=b[i]
					}
				}
			}
		}
		
		if (max_max==1) {
			fmt.Println("-1")
		} else {
			fmt.Printf("%d ", max_max)
		}
		fmt.Printf("\n")
	}
} 