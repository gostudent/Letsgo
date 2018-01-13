---
published: true
author: surasnayak
layout: post
title: Find out the Team Solution in Go
---

## Find out the Team

### Question:

A school is conducting a debate competition in which participation is in teams. Stefan is the manager of the event. So has to allot students to different teams. There are 400 students with roll numbers from 1 to 400. Stefan thinks that he is smart so he used a logic to divide students in teams using their roll numbers. He has challenged you to find out the logic behind it. One team can have a maximum 3 members. Stefan has divided students using their roll numbers :
ROLL NO. OF STUDENT TEAM NO.
1 1
2 1
.............. ............

79 39
80 40

### Constraints :

1 < = T < = 400
1 < = X < = 1000

### Input

The first line of the input contains an integer T denoting the number of test cases. The description of each testcase follows. Each test case contains a single line with a single integer X ( roll number of the student ) .

### Output

For each test case output a new line with a single integer which is the team number .

### Sample TestCase Input :

- 5
- 4
- 89
- 90
- 897
- 435

### Sample TestCase Output :

- 2
- 44
- 45
- 448
- 217


## Solution

- **Language Used:** go

```go
package main
import "fmt"

func main(){
	
	 var t int
	 fmt.Scan(&t)
	 for i:=0; i<t; i++ {
	 	var n int
	 	fmt.Scan(&n)
	 	if(n == 1) {
	 		fmt.Println("1")
	 	} else if(n%2 == 0) {
	 		fmt.Println(n/2)
	 	} else if(n%2 != 0) {
	 		fmt.Println((n-1)/2)
	 	}
	 }
}
```
