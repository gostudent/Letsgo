package main
 
import "fmt"
 
func main() {
    var n int
    fmt.Scan(&n)
    ar := make([]int, n)
    ReadN(ar, 0, n)
 
    for j := 0; j < n; j++ {
        r := 1
        s := 1
        for k := 0; k < n; k++ {
            if (j != k) && (ar[k] < ar[j]) {
                r += 1
            }
            if (j != k) && (ar[k] == ar[j]) {
                s += 1
            }
        }
        fmt.Print(float32(float32(r) + float32(s - 1)/(float32(2))));
        fmt.Print(` `)
    }
}
 
func ReadN(all []int, i, n int) {
    if n == 0 {
        return
    }
    fmt.Scan(&all[i])
    ReadN(all, i+1, n-1)
}
 