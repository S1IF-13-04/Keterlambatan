package main
import "fmt"
func main(){
	var x,y int64

	fmt.Scan(&x,&y)
	if x>=5000 && y >=24{
		fmt.Print("VVIP Backstage")
	}else if x>=2000 && y>=12 {
		fmt.Print("VIP Soundcheck")
	}else if x>500 {
		fmt.Print("Festival Ground")
	}else{
		fmt.Print("Menonton dari YouTube")
	}
}