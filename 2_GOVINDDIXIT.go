package main

import "fmt"

func main() {
        var z int
        fmt.Scan(&z)
        for j := 0; j < z; j++ {
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
            var maximum int
            var minimum int
            if x>=y {
                maximum =x
                minimum =y
            }
            if y>x {
                maximum =y
                minimum =x
            }
            answer :=(maximum/minimum)
            fmt.Printf("%v\n",answer);
        }
}
