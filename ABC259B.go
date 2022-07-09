package main
import "fmt"
import "math"

func main(){
	var a,b,d float64
	fmt.Scan(&a,&b,&d)

	var x = a*math.Cos(d*math.Pi/180) - b*math.Sin(d*math.Pi/180)
	var y = a*math.Sin(d*math.Pi/180) + b*math.Cos(d*math.Pi/180)
	fmt.Println(x,y)
	
}
