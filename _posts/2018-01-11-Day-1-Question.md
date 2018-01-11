---
published: true
---
# Day 1
# Issue 1

## **Question:**
## **Given three positive integers X, Y and P. Here P denotes the number of turns. Whenever the turn is odd X is multiplied by 2 and in every even turn Y is multiplied by 2. The task is to find the value of max(X, Y) ÷ min(X, Y) after the complete P turns.**

## **Input:**
## **The first line of the input contains a single integer T, denoting the number of test cases. Then T test case follows, a single line of the input containing three integers X,Y and P.**

## **Constraints:**
## **1<=T<=100**
## **1<=X<=Y<=P<=1000**

## **One of the solution:**
	package main
	import "fmt"
	func main() {
        var t int
        fmt.Scan(&t)
        for j := 0; j < t; j++ {
            var x int
            fmt.Scan(&x)
            var y int
            fmt.Scan(&y)
            var p int
            fmt.Scan(&p)
            for i := 1; i <= p; i++ {
                if i%2!=0 {
                    x=x*2
                }
                if i%2==0 {
                    y=y*2
                }
            }
            var max int
            var min int
            if x>=y {
                max =x
                min =y
            }
            if y>x {
                max =y
                min =x
            }
            ans :=(max/min)
            fmt.Printf("%v\n",ans);
        }
}
