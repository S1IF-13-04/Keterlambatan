package main

import "fmt"

func main() {
	var n, x, yes int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		if x > 75 {
			fmt.Println("Juri bilang YES")
			yes++
		} else {
			fmt.Println("Juri Bilang NO")
		}
	}
	fmt.Println("Toal YES", yes)
}
