//2次元スライス

package main
import "fmt"
import "bufio"
import "os"

func main(){
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

	var R,C int
	fmt.Fscan(in, &R,&C)
	
	N := 15
	done := make([][]bool, N)
	for i := 0; i < N; i++ {
	    done[i] = make([]bool, N)
	}
	
	res := make([][]string, N)
	for i := 0; i < N; i++ {
	    res[i] = make([]string, N)
	}
	
	for k := 0; k < 8; k++ {
	    for y := -k; y < k+1; y++ {
	        for x := -k; x < k+1; x++ {
	            if done[y+7][x+7] == true {
	                continue
	            }
	            if k%2 == 0 {
	                res[y+7][x+7] = "white"
	            } else {
	                res[y+7][x+7] = "black"
	            }
	            done[y+7][x+7] = true
	        }
	    }
	}
    fmt.Fprintln(out, res[R-1][C-1])
    
//	for i := 0; i < N; i++ {
//	    fmt.Fprintln(out, res[i])
//	}
}

