package main
import "fmt"

func main(){
	var tc int
	fmt.Scan(&tc)
	t:=0
	for t<tc{
	   var x int
	   fmt.Scan(&x)
	   y :=x/2
	   fmt.Println(y)
	   t+=1
	}
}