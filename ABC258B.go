// dxとdyで周囲を確認

package main
import "fmt"
//import "bufio"
//import "os"

func main(){
    var N int
    var ans int64
    fmt.Scan(&N)
    dx := [...]int{-1,-1,-1,0,0,1,1,1}
    dy := [...]int{-1,0,1,-1,1,-1,0,1}
    
    A := make([][]int64, N)
    for i := 0; i < N; i++ {
        A[i] = make([]int64, N)
        var a string
        fmt.Scan(&a)
        for j := 0; j < N; j++ {
            A[i][j] = int64(a[j] - '0')
        }
    }
    
    for y := 0; y < N; y++ {
        for x := 0; x < N; x++ {
            for d := 0; d < 8; d++ {
                var tmp int64
                for i := 0; i < N; i++{
                    tmp *= 10
                    var Y,X int
                    Y = (y + dy[d]*i)%N
                    if Y < 0 {
                        Y += N
                    }
                    X = (x + dx[d]*i)%N
                    if X < 0 {
                        X += N
                    }
                    
                    tmp += A[Y][X]
//                    fmt.Println(tmp,Y,X)
                }
                if ans < tmp {
                    ans = tmp
                }
            }
        }
    }
    fmt.Println(ans)
}

