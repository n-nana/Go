package main
import "fmt"

func main(){
	var N,X,Y,Z,a,b int
	var A,B []int //slice
	
	fmt.Scan(&N,&X,&Y,&Z)
	
	for i := 0; i < N; i++ {
	    fmt.Scan(&a)
	    A = append(A,a)
	}

	for i := 0; i < N; i++ {
	    fmt.Scan(&b)
	    B = append(B,b)
	}
	
	flag := make([] bool, N)
	
	// Mathmatics
    for i := 0; i < X; i++ {
        mx := -1
        idx := -1
        for j := 0; j < N; j++ {
            if flag[j] {
                continue
            }
            if A[j] > mx {
                mx = A[j]
                idx = j
            }
        }
        flag[idx] = true
    
    }
//    fmt.Println(flag)

    // English
    for i := 0; i < Y; i++ {
        mx := -1
        idx := -1
        for j := 0; j < N; j++ {
            if flag[j] {
                continue
            }
            if B[j] > mx {
                mx = B[j]
                idx = j
            }
        }
        flag[idx] = true
    }
//	fmt.Println(flag)

    // Mathmatics + English
	for i := 0; i < Z; i++ {
        mx := -1
        idx := -1
        for j := 0; j < N; j++ {
            if flag[j] {
                continue
            }
            if A[j]+B[j] > mx {
                mx = A[j]+B[j]
                idx = j
            }
        }
        flag[idx] = true
    }
    
    for i := 0; i <N; i++ {
        if flag[i] == true {
            fmt.Println(i+1)
        }
    }
//	fmt.Println(flag)
}
