//min, float64

package main
import "fmt"
import "math"

func main(){
    var N,a int
    P := make([]int, 5)
    fmt.Scan(&N)

    for i := 0; i<N; i++{
        fmt.Scan(&a)
        P[0]++
        for j := 0; j<4; j++{
            var crr = 3-j
            var val = float64(crr+a)
            var nxt = int(math.Min(val,4))
            P[nxt] += P[crr]
            P[crr] = 0
        }
    }
    fmt.Println(P[4])
}

