---
published: true
author: marshallKR7
layout: post
title:  Reverse words in a given String Solution in Go
---

## Reverse words in a given String

### Question:

Given a String of length N reverse the words in it. Words are separated by dots.


### Constraints :

- 1<=T<=20
- 1<=Length of String<=2000

### Sample Input:

The first line contains T denoting the number of testcases. Then follows description of testcases. Each case contains a string containing spaces and characters.

### Sample Output:

For each test case, output a single line containing the reversed String.


- Sample TestCase
- Input :
- 2
- i.like.this.program.very.much
- pqr.mno


### Output :

- much.very.program.this.like.i
- mno.pqr



## Solution

- **Language Used:** go

```go
    package main
     
    import (
    	"fmt"
    	"strings"
    )
     
    func reverse_words(s string) string {
    	words := strings.Fields(s)
    	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
    		words[i], words[j] = words[j], words[i]
    	}
    	return strings.Join(words, ".")
    }
     
    func main() {
    	var t int
    	fmt.Scan(&t)
    	for j := 0; j < t; j++ {
    		var a string
    		fmt.Scan(&a)
    		fmt.Println(reverse_words(strings.Replace(a, ".", " ", -1)))
    	}
    }
  ```
