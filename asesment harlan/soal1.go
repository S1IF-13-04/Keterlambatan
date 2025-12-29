package main

import "fmt"

func main() {
	var pengeluaran, lamajoin int

	fmt.Scan(&pengeluaran, &lamajoin)

	if pengeluaran >= 5000 && lamajoin >= 24 {
		fmt.Println("VVIP Backstage")
	} else if pengeluaran >= 2000 && lamajoin >= 12 {
		fmt.Println("VIP Soundcheck")
	} else if pengeluaran >= 500 {
		fmt.Println("Festival Ground")
	} else {
		fmt.Println("Menonton dari YouTube")
	}
}
