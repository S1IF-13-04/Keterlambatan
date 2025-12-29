package main
import "fmt"

func main(){
	var x,y,i int64
	fmt.Scan(&x)
	fmt.Println("masukan angka : ",x)
	for i=x; i>=1;i=x{
		y+= x%10
		x= x/10
	} 
	fmt.Print("hasil penjumlahan :",y)
}