    package main
     
    import "fmt"
     
    func main() {
    	var testcase int
    	fmt.Scan(&testcase)
    	for j := 0; j < testcase; j++ {
    		var n int
    		fmt.Scan(&n)
    		answer := n/2
    		fmt.Println(answer)
    	}
    }
