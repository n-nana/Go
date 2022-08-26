package main
import "fmt"
import "bufio"
import "os"

//import "reflect"
//import "sort"
//import "strconv"
//import "strings"

func main(){
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

	var N int
	fmt.Fscan(in, &N)
    
// 1D_slice_string
//    A := make([]string, N)
//    for i := 0; i < N; i++ {
//        var a string
//        fmt.Fscan(in, &a)
//        A[i] = a
//    }
    
//    2D_slice_bool
//    A := make([][]bool, N)
//    for i := 0; i < N; i++ {
//        A[i] = make([]bool, N)
//    }
    
    fmt.Fprintln(out, N)
}
