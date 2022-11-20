//ABC278E - Grid Filling
//2次元累積和


package main
import "fmt"
import "bufio"
import "os"

//import "reflect"
//import "sort"
//import "strconv"
//import "strings"

const MOD1 int = 998244353
var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func main(){
    defer out.Flush()

	var H,W,N,h,w int
	fmt.Fscan(in, &H,&W,&N,&h,&w)
    
    // 標準入力(2次元リスト)
    A := make([][]int, H)
    for i := 0; i < H; i++ {
        A[i] = make([]int, W)
        for j := 0; j < W; j++ {
            var a int
            fmt.Fscan(in, &a)
            A[i][j] = a
        }
    }
    
    // 空の3次元リスト
    RU := make([][][]int, N+1)
    for p := 0; p < N+1; p++ {
        RU[p] = make([][]int, H+3)
        for i := 0; i < H+3; i++ {
            RU[p][i] = make([]int, W+3)
            for j := 0; j < W+3; j++ {
                RU[p][i][j] = 0
            }
        }
    }
  
    // 2次元累積和
    for p := 1; p < N+1; p++ {
        for i := 0; i < H; i++ {
            for j := 0; j < W; j++ {
                RU[p][i+1][j+1] = RU[p][i+1][j] + RU[p][i][j+1] - RU[p][i][j]
                if p == A[i][j] {
                    RU[p][i+1][j+1] ++
                }
            }
        }
    }
    
    K := H-h+1
    L := W-w+1
    
    // 空の2次元リスト
    res := make([][]int, K)
    for i := 0; i < K; i++ {
        res[i] = make([]int, L)
        for j := 0; j < L; j++ {
            res[i][j] = 0
        }
    }
    
    for i := 0; i < K; i++ {
        for j := 0; j < L; j++ {
            for p := 1; p < N+1; p++ {
                cnt := RU[p][H][W] - RU[p][i+h][j+w] + RU[p][H][j] + RU[p][i][W]
                if cnt > 0 {
                    res[i][j] ++
                }
            }
        }
    }

    for i := 0; i < K; i++ {
        for j := 0; j < L; j++ {
            // 空白区切りで出力
            fmt.Fprint(out,res[i][j])
            fmt.Fprint(out," ")
        }
        fmt.Fprintln(out) //改行
    }
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
