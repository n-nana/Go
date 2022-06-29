package main
import "fmt"
import "bufio"
import "sort"
import "os"

func main(){
    in := bufio.NewReader(os.Stdin)
//    out := bufio.NewWriter(os.Stdout)
//    defer out.Flush()
    
    var N int
    fmt.Fscan(in, &N)
    D := make([] int, N)
    
    for i := 0; i < N; i++ {
        fmt.Fscan(in, &D[i])
    }
    sort.Ints(D)
    fmt.Println(D[N/2]-D[N/2-1])
}

