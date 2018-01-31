---
published: true
author: ADESH SHUKLA
layout: post
title: Problem
---

## Question :

No problem statement.Find the logic from the given sample input/output.
And answer Q queries.

## Sample TestCase :

### Input :

8
10
30
45
9
69
77
127
150

### Output :

8
42
33
4
27
19
1
222

## Solution :

```go

package main
import ("fmt")

func main() {
  var n int
  fmt.Scan(&n)
  for i := 0; i < n; i++ {

      var t int
      fmt.Scan(&t)
      var su int = 0
      for i := 1; i < t; i++ {

        if(t%i==0){
            su=su+i
    }

  }

  fmt.Println(su)
  }

 }

``