package main
import "fmt"
func main(){
    
    var N,K,Q,a,l,p int
    blk := " "
    fmt.Scan(&N,&K,&Q)
    var A[]int

    for i := 0; i < K; i++{
        fmt.Scan(&a)
        A = append(A,a)
    }
    for i := 0; i < Q; i++{
        fmt.Scan(&l)
        p = l - 1
        if A[p] == N{
            continue
        } else{
            if p == K-1{
                A[p] ++
            } else if A[p+1] == A[p] + 1{
                continue
            } else{
                A[p]++
            }
        }
    }
    for i := 0; i < K; i++{
        fmt.Print(A[i])
        fmt.Print(blk)
    }
}

