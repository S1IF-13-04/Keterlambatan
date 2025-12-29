package main
import "fmt"

func main()  {
	var x,y,t,h int64
	fmt.Scan(&x)

	for t=0 ; t< x; t=t{
		fmt.Scan(&y)
		t+=y
		if t <= x{	
			h=t
			fmt.Printf("Masuk %d liter. Total: %d\n",y,h)
		}else {
			fmt.Printf("Luber! Total Tetap %d Tidak bisa tambah %d\n",h,y)
		}
	}
}