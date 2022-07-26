//入出力高速化

package main
import "fmt"
import "bufio"
import "os"

func main(){
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

	var flag bool
	var N int
	fmt.Fscan(in, &N)

	A := make([]string, N)
	for i:= 0; i < N; i++ {
	    var a string
	    fmt.Fscan(in, &a)
	    A[i] = a
	}
	
	for i := 0; i < N; i++ {
	    for j := 0; j < N; j++ {
	        if string(A[i][j]) == "W" && string(A[j][i]) != "L" {
	            flag = true
	        } else if string(A[i][j]) == "L" && string(A[j][i]) != "W" {
	            flag = true
	        } else if string(A[i][j]) == "D" && string(A[j][i]) != "D" {
	            flag = true
	        }
	    }
	}
	if flag == false {
	    fmt.Fprint(out,"correct")
	} else {
	    fmt.Fprint(out,"incorrect")
	}
}


