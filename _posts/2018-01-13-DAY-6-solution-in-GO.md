---
published: true
---
# Solution for 'Set all the bits in given range of a number' in GO. 

**Author@** Suras Kumar Nayak

## Set all the bits in given range of a number

### Question:

Given a non-negative number N and two values L and R. The problem is to set all the bits in the range L to R in the binary representation of N.

### Constraints :

- 1<=T<=100
- 1<=N<=232
- 1<=L<=R<=32

### Input:

The first line of input contains an integer T denoting the number of test cases. Then T test cases follow. Each test case contains three integers N, L and R as input.

### Output:

For each test case, set the bits in given range and print the modified number in new line.

### Sample Test Case:

### Input :

- 2
- 17 2 3
- 8 1 2

### Output :

- 23
- 11

### Explanation:

- Input : N = 17, L = 2, R = 3
- Output : 23
- (17)10 = (10001)2
- (23)10 = (10111)2
- The bits in the range 2 to 3 in the binary
representation of 17 are set.

## Solution

- **Language Used:** go

```go
package main
import "fmt"
 
func main(){
	var t int
	fmt.Scan(&t)
	
	for j := 1; j <= t; j++ {
		
		var n int
		fmt.Scan(&n)
		
		var l int
		fmt.Scan(&l)
		
		var r int
		fmt.Scan(&r)
		
		fmt.Println(n|(((1 << uint(l - 1)) - 1) ^ ((1 << uint(r)) - 1))) 
	}
}

```
