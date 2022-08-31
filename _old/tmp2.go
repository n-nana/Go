package main
import "fmt"
import "bufio"
import "os"
import "math"

const INF int = 1<<60

func main(){
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

    for {
    	var N,M int
	    fmt.Fscan(in, &N,&M)
	    
	    if N == 0 && M == 0 {
	        break
	    } else {
	        
	        C := make([]int, M)
	        for k := 0; k < M; k++ {
	            var c int
	            fmt.Fscan(in, &c)
	            C[k] = c
            }
            
            X := make([]int, N)
	        for i := 0; i < N; i++ {
	            var x int
	            fmt.Fscan(in, &x)
	            X[i] = x
            }
//            fmt.Fprintln(out, C)
//            fmt.Fprintln(out, X)
            
            dp := make([][]int, 2)
            for i := 0; i < 2; i++ {
                dp[i] = make([]int, 256)
                for j := 0; j < 256; j++ {
                    dp[i][j] = INF
                }
            }
            dp[0][128] = 0
            
            for i := 0; i < N; i++ {
                x := X[i]
                for j := 0; j < 256; j++ {
                    if dp[0][j] == INF {
                        continue
                    } else {
                        for k := 0; k < M; k++ {
                            nxt := j + C[k]
                            if nxt < 0 {
                                nxt = 0
                            } else if nxt > 255 {
                                nxt = 255
                            }
                            cal := int(math.Pow(float64(x - nxt), 2))
//                            fmt.Fprintln(out, dp[i][j] + cal)
                            if dp[1][nxt] > dp[0][j] + cal {
                                dp[1][nxt] = dp[0][j] + cal
                            }
                        }
                    }
                }
                for j := 0; j < 256; j++ {
                    dp[0][j] = dp[1][j]
                    dp[1][j] = INF
                }
//                fmt.Fprintln(out, dp)
//                fmt.Fprintln(out, "")
//                dp[0] = dp[1]
            }
//            dp[0] = dp[1]

            res := INF
            for j := 0; j < 256; j ++ {
                if dp[0][j] < res {
                    res = dp[0][j]
                }
            }
            fmt.Fprintln(out, res)
	    }
    }
}
