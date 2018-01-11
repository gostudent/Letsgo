package main

import "fmt"

func main() {
    
    var t int
    fmt.Scan(&t)
    
    for j := 0; j < t; j++ {
        var x, y, p int
        fmt.Scan(&x)
        fmt.Scan(&y)
        fmt.Scan(&p)
        for i := 1; i <= p; i++ {
            if i % 2 != 0 {
                x *= 2
            }
            if i % 2 == 0 {
                y *= 2
            }
        }
        var max, min int
        if x >= y {
            max = x
            min = y
        }
        if y > x {
            max = y
            min = x
        }
        m := (max/min)
        fmt.Println(m);
    }
}