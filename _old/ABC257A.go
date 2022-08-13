package main
import "fmt"
func main(){
    var N,X,k int
    fmt.Scan(&N,&X)
    P := make([]string, N*26)
k = 0
    for i := 65; i<91; i++{
        for j := 0; j<N; j++{
            P[k] = string(i)
            k++
        }
    }
    fmt.Println(P[X-1])
//    fmt.Println(P)
}
