package main 
import "fmt"

func main(){
	var x,i,y,j int64
	fmt.Print("Masukan hari ke 1 sampai hari ke 2: ")
	fmt.Scan(&x,&y)
	j = 1
	for i = x;i<=y;i++{
		j *= i
	}
	fmt.Print(j)
}