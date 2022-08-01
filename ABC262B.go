package main
import "fmt"
import "bufio"
import "os"

func main(){
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

	var N,M int
	var U,V int
	
	fmt.Fscan(in, &N, &M)
    
    G := make([][]bool, N)
    for i := 0; i < N; i++ {
        G[i] = make([]bool, N)
    }
    
    for j := 0; j < M; j++ {
        fmt.Fscan(in, &U, &V)
//        fmt.Fprintln(out, U, V)
        G[U-1][V-1] = true
        G[V-1][U-1] = true
    }
//    fmt.Fprintln(out, G)

    res := 0
    for a := 0; a < N-2; a++ {
        for b := a+1; b < N-1; b++ {
            for c := b+1; c < N; c++ {
                if G[a][b] == true && G[b][c] == true && G[c][a] == true {
                    res++
                }
            }
        }
    }
    fmt.Fprintln(out, res)
}


