package main
import "fmt"
import "bufio"
import "os"

//import "strings"
//import "sort"
//import "reflect"

func main(){
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

	var N int
	fmt.Fscan(in, &N)
    
// 1D_slice
//    A := make([]string, N)
//    for i := 0; i < N; i++ {
//        var a string
//        fmt.Fscan(in, &a)
//        A[i] = a
//    }
    
    fmt.Fprintln(out, N)
}

