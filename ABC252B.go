package main
import "fmt"

func main(){
	var N,K,a,b int
	var mx int = -1
	var A,B,idx []int //slice
	var ans string = "No"

	fmt.Scan(&N,&K)

// read A
	for i := 0; i < N; i++{
	    fmt.Scan(&a)
	    A = append(A,a)
	}
	
// max. value of A
	for i := 0; i < N; i++{
	    if A[i] > mx{
	        mx = A[i]
	    }
	}
	
// index of favorite item
	for i := 0; i < N; i++{
	    if A[i] == mx{
	        idx = append(idx,i)
	    }
	}
	
// read B
	for i := 0; i < K; i++{
	    fmt.Scan(&b)
	    B = append(B,b)
	}
	
// check
	for i := 0; i < len(idx); i++{
	    for j := 0; j < K; j++{
	        if idx[i] == B[j]{
	            ans = "Yes"
	        }
	    }
	}
	
	fmt.Println(ans)

}

