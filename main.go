package main

import (
	"fmt"
	"homework/misollar"
	"os"
)

func main(){
	var n int
	for {
		fmt.Printf("\n1.masala-1\n2.masala-2\n3.chiqish\n")
		fmt.Scan(&n)
		switch n {
			case 1:
				misollar.Masala1()
			case 2:
				misollar.Masala2()
			case 3:
				os.Exit(0)
			default:
				fmt.Printf("invalid input\n")
		}
	}
}