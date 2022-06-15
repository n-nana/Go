package main
import "fmt"

func main(){
	var a,b,c int
	fmt.Scan(&a, &b, &c)
	
	if (a <= b && b <= c) || (c <= b && b <= a){
	    fmt.Print("Yes")
	} else{
	    fmt.Print("No")
	}
}

