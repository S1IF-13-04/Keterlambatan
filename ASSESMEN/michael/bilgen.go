package main 
import "fmt"

func main(){
	var x,i,y int64
	y = 2
	fmt.Print("Masukan jumlah bilangan genap: ")
	fmt.Scan(&x)
	for i=1;i<=x;i++{
		fmt.Println(y)
		y +=2
	}
}