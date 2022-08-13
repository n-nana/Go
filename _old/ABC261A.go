package main
import "fmt"

func main(){
	var L1,R1,L2,R2 int
	fmt.Scan(&L1,&R1,&L2,&R2)
	k := 0
  
	if R1 < R2 {
	    k = R1
	} else {
	    k = R2
	}
  
	if L1 < L2 {
	    k -= L2
	} else {
	    k -= L1
	}
  
	if k > 0 {
	    fmt.Println(k)
	} else {
	    fmt.Println(0)
	}
	
}

