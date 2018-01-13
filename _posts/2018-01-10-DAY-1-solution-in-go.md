---
published: true
author: nikki1016
layout: post
title: Even-Odd turn game Solution in Go
---

## Even-Odd turn game

### Question:

Given three positive integers X, Y and P. Here P denotes the number of turns. Whenever the turn is odd X is multiplied by 2 and in every even turn Y is multiplied by 2. The task is to find the value of max(X, Y) ÷ min(X, Y) after the complete P turns.


### Constraints :

1<=T<=100
1<=X<=Y<=P<=1000

### Sample Input:

The first line of the input contains a single integer T, denoting the number of test cases. Then T test case follows, a single line of the input containing three integers X,Y and P.

### Sample Output:

For each test case, print the value of max(X, Y) ÷ min(X, Y) after the complete P turns .


- Sample TestCase
- Input :
- 2
- 1 2 1
- 3 7 2


### Output :

- 1
- 2



## Solution

- **Language Used:** go

```go
package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)
	for ; t > 0 ; t-- {
		var x, y, p int;
		var mx, mn int = 1,2;
		fmt.Scan(&x, &y, &p);
		for p > 0{
			if( p %2 ==1 ) {
                		x = x*2
           		} else {
              			y = y*2
           		}
			p = p-1
		}
		if x > y{
			mx = x;
			mn = y;
		} else{
		    mx = y
		    mn = x
		}
	
		var ans int = mx/mn
		fmt.Println(ans)
	}	
}
```

