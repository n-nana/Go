// int > str

package main
import "fmt"
import "strconv"
// strconv.Itoa(int), int -> str
// strconv.Atoi(str), str -> int

func main(){
	var N,a,b int
	fmt.Scan(&N)
    a = N%10
    b = N/10%10
    A := strconv.Itoa(a)
    B := strconv.Itoa(b)

//    fmt.Println(b,a)
    fmt.Print(B,A)
}


