package main
import "fmt"

func main(){
	var S string
	fmt.Scan(&S)
	
	A := make([] int, 26) // bucket
	for i := 0; i < 3; i++ {
	    A[S[i]-97] += 1
	}
  
	ans := -1
	for i := 0; i < 26; i++ {
	    if A[i] == 1 {
	        ans = i
	    }
	}
	if ans == -1 {
	    fmt.Println(ans)
	} else {
	    fmt.Println(string(ans+97)) // unicode変換
	}

}
