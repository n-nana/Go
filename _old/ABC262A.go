package main
import "fmt"
import "bufio"
import "os"

func main(){
    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

	var Y int
	fmt.Fscan(in, &Y)
	rem := Y%4
	
    fmt.Fprint(out,Y+(6-rem)%4)
	
}


