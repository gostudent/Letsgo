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
