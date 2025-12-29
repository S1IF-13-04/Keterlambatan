package main
import "fmt"

func main(){
	var x,i,j int64
	fmt.Scan(&x)

	for i=1;i<=x;i++{
		for j=1;j<=i;j++{
			fmt.Printf("%d ",i)
		}
		fmt.Println()
	}
}