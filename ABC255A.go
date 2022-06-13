package main
import "fmt"

func main(){
	var r,c,p1,p2,p3,p4 int
	fmt.Scan(&r, &c)
	fmt.Scan(&p1, &p2)
	fmt.Scan(&p3, &p4)
	if r == 1{
		if c == 1{
			fmt.Println(p1)
		} else if c == 2{
			fmt.Println(p2)
		}
	} else if r == 2{
		if c == 1{
			fmt.Println(p3)
		} else if c == 2{
			fmt.Println(p4)
		}
	}
}

