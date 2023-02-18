// ABC075D - Axis-Parallel Rectangle

package main
import "fmt"
import "bufio"
import "os"
import "math"

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main(){
    defer out.Flush()

	var N,K int
	fmt.Fscan(in, &N,&K)

	A := make([][]int, N)
    for i := 0; i < N; i++ {
        A[i] = make([]int, 2)
        for j := 0; j < 2; j++ {
            var v int
            fmt.Fscan(in, &v)
            A[i][j] = v
        }
    }
    
    res := 1<<62
    for p1 := 0; p1 < N; p1++ {
        for p2 := 0; p2 < N; p2++ {
            for p3 := 0; p3 < N; p3++ {
                for p4 := 0; p4 < N; p4++ {
                    l := min(A[p1][0], A[p2][0])
                    r := max(A[p1][0], A[p2][0])
                    u := min(A[p1][1], A[p2][1])
                    d := max(A[p1][1], A[p2][1])
                    
                    l = min(A[p3][0], l)
                    r = max(A[p3][0], r)
                    u = min(A[p3][1], u)
                    d = max(A[p3][1], d)
                    
                    l = min(A[p4][0], l)
                    r = max(A[p4][0], r)
                    u = min(A[p4][1], u)
                    d = max(A[p4][1], d)
                    
                    cnt := 0
                    for i := 0; i < N; i++ {
                        var x,y int
                        x = A[i][0]
                        y = A[i][1]
                        if (l <= x) && (x <= r) && (u <= y) && (y <= d) {
                            cnt++
                        }
                    }
                    if cnt >= K {
                        tmp := (r-l)*(d-u)
                        res = min(tmp, res)
                    }
                }
            }
        }
    }
    fmt.Fprintln(out, res)
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

func pow(p,q int) int {
    return int(math.Pow(float64(p), float64(q)))
}
