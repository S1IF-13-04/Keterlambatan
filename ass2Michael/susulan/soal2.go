package main
import "fmt"

func main(){
	var x,y,i,t int64

	fmt.Scan(&x)
	t=0
	for i=1;i<=x;i++{
		fmt.Scan(&y)
		if y > 75{
			fmt.Println("Juri bilang: Yes")
			t+=1
		} else {
			fmt.Println("Juri bilang: No")
		}
	}
	fmt.Printf("Total Yes: %d",t)
}