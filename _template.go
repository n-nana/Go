package main
import "fmt"
import "bufio"
import "os"

//import "reflect"
//import "sort"
//import "strconv"
//import "strings"

const MOD1 int = 998244353
var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main(){
    defer out.Flush()

	var N,M int
	fmt.Fscan(in, &N,&M)
    
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
    
    fmt.Fprintln(out, N,M,max(N,M),min(N,M),abs(N),abs(M))
}


//-------------------------------------------------------------
func max(a,b int) int {
    if a < b {
        return b
    } else {
        return a
    }
}

func min(a,b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}

func abs(a int) int {
    if a < 0 {
        return -a
    } else {
        return a
    }
}

