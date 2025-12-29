package main
import "fmt"

func main(){
	var x,y,t,h int64
	var p string

	fmt.Println("==== DAFTAR PRODUK TOKO MAS RUSDI ====")
	fmt.Println("1. Si Imut                  - Rp15.000")
	fmt.Println("2. Ambatron                 - Rp250.000")
	fmt.Println("3. Mas Amba                 - Rp150.000")
	fmt.Println("4. Abatukam                 - Rp10.000")
	fmt.Print("Pilih Produk (1-4) : ")
	fmt.Scan(&x)
	fmt.Print("Masukan Jumlah Beli : ")
	fmt.Scan(&y)

	if x == 1 {
		h = 15000
		t = h*y
		p = "Si Imut"
	}else if x == 2{
		h = 250000
		t = y *h
		p = "Ambatron"
	}else if x == 3{
		h = 150000
		t = y *h
		p = "Mas Amba"
	}else if x == 4{
		h = 10000
		t = y *h
		p = "Abatukam"
	}else{
		fmt.Print("Tolong Pilih dari 1-4")
	}
	fmt.Println("=== STRUK PEMBAYARAN ===")
	fmt.Println("Produk: ",p)
	fmt.Println("Jumlah: ",y)
	fmt.Println("Harga : ",h)
	fmt.Println("Total : ",t)
	
}