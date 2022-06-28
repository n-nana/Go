//ほとんど同じだが、、

package main
import "fmt"
import "bufio"
import "os"
//import "math"

const INF = 1 << 60

func main(){
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

    var U,V,N,M,ans int

//    fmt.Scan(&N,&M)
    fmt.Fscan(in, &N,&M)
    
    E := make([][]int, N+1)

    for i := 0; i<M; i++ {
        fmt.Fscan(in, &U,&V)
        E[U] = append(E[U],V)
        E[V] = append(E[V],U)
    }

    res1 := bfs(1,N,E)
    res2 := bfs(N,N,E)

    for i := 1; i<N+1; i++ {
        ans = res1[N]
        if res1[0] + res2[i] < ans {
            ans = res1[0] + res2[i]
        }
        if res2[0] + res1[i] < ans {
            ans = res2[0] + res1[i]
        }
        if ans == INF {
            fmt.Fprint(out, -1)
        } else {
            fmt.Fprint(out, ans)
        }
        fmt.Fprint(out, " ")
    }
}    


func bfs(s,N int, E [][]int) ([]int) {
    var crr,nxt int

    dist := make([]int, N+1)
    for i := 0; i<N+1; i++ {
        dist[i] = INF
    }
    dist[s] = 0

    que := make([]int, 0)
    que = append(que,s)
    
    for len(que) > 0 {
        crr = que[0]
        que = que[1:]
        for i := 0; i<len(E[crr]); i++ {
            nxt = E[crr][i]
            if dist[nxt] == INF {
                dist[nxt] = dist[crr] + 1
                que = append(que, nxt)
            }
        }
    }
    return dist
}
