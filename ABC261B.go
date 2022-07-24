package main
import "fmt"

func main(){
	var N int
	var flag bool
	
	fmt.Scan(&N)
	
	A := make([]string, N)
	for i:= 0; i < N; i++ {
	    var a string
	    fmt.Scan(&a)
	    A[i] = a
	}
	
	for i := 0; i < N; i++ {
	    for j := 0; j < N; j++ {
//	        println(string(A[i][j]))
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
	    fmt.Println("correct")
	} else {
	    fmt.Println("incorrect")
	}
}

