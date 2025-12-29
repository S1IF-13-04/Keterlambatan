package main

import "fmt"

func main() {
	var n, jumlah int
	fmt.Scan(&n)

	for n > 0 {
		jumlah += n % 10
		n /= 10
	}
	fmt.Println(jumlah)
}
