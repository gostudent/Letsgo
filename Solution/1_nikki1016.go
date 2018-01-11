package main

import (
	"fmt"
)

func main() {
	var t int
	_, err := fmt.Scan(&t)
	  if err != nil {
	      fmt.Println(24444)
	      return
	  }
	for t > 0{
		var x, y, p int;
		var mx, mn int = 1,2;
		_, err := fmt.Scan(&x, &y, &p);
		if err != nil {
	      fmt.Println(2222)
	      return
	  }
		for p > 0{
			if( p %2 ==1 ) {
                x = x*2
           } else {
              y = y*2
           }
			p = p-1
		}
		t = t-1;
		if x > y{
			mx = x;
			mn = y;
		} else{
		    mx = y
		    mn = x
		}
	
		var ans int = mx/mn
		fmt.Println(ans)
	}	
}
