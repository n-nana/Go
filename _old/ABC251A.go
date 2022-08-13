package main
import "fmt"

func main(){
	var S,res string
	fmt.Scan(&S)
	for len(res) < 6{
	    res += S
	}
    fmt.Print(res)
}

