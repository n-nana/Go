//べき乗（math.Pow）, math.Powはfloat64型
//bit全探索
//スライスの比較（reflect.DeepEqualでスライスを比較すると遅いかも？）

package main
import "fmt"
import "bufio"
import "os"
import "math"
//import "reflect"

func main(){
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

	var H1,W1 int
	fmt.Fscan(in, &H1,&W1)
	A := make([][]int, H1)
	for i := 0; i < H1; i++ {
	    A[i] = make([]int, W1)
	    for j := 0; j < W1; j++ {
	        var a int
	        fmt.Fscan(in, &a)
	        A[i][j] = a
	    }
	}

	var H2,W2 int
	fmt.Fscan(in, &H2,&W2)
	B := make([][]int, H2)
	for i := 0; i < H2; i++ {
	    B[i] = make([]int, W2)
	    for j := 0; j < W2; j++ {
	        var b int
	        fmt.Fscan(in, &b)
	        B[i][j] = b
	    }
	}
	
	ans := "No"
    for i := 1; i < int(math.Pow(2,float64(H1))); i++ {
        Y := make([]int, 0)
        for y := 0; y < H1; y++ {
            if (i>>y)&1 == 1 {
                Y = append(Y,y)
            }
        }
        
        for j := 1; j < int(math.Pow(2,float64(W1))); j++ {
            X := make([]int, 0)
            for x := 0; x < W1; x++ {
                if (j>>x)&1 == 1{
                    X = append(X,x)
                }
            }
            
            if ! (len(Y) == H2 && len(X) == W2) {
                continue
            }
            
            flag := true
            for h := 0; h < len(Y); h++ {
                for w := 0; w < len(X); w++ {
                    if B[h][w] != A[Y[h]][X[w]]{
                        flag = false
                    }
                }
            }
            if flag == true {
                ans = "Yes"
            }
        }
    }
    fmt.Fprintln(out, ans)
}


