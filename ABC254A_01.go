// strをslice

package main
import "fmt"
//import "strconv"
// strconv.Itoa(int), int -> str
// strconv.Atoi(str), str -> int

func main(){
	var N string
	fmt.Scan(&N)
    fmt.Print(N[1:])
}


