package main 
import "fmt"

func main(){
	var x,p,k,i,e int64
	fmt.Print("Masukan uang dalam keping: ")
	fmt.Scan(&x)
	
	p = x/800
	k = x % 800
	ke := k/ 80
	i = x %80
	ik:= i / 8
	e = x%8

	fmt.Printf("%d peti, %d karung, %d ikat, dan %d keping",p,ke,ik,e)

}