package misollar

import (
	"fmt"
	"homework/methods/factorial"
)

func Masala1(){
	var n int
	ch :=  make(chan int)

	fmt.Print("enter a number: ")
	fmt.Scan(&n)
	go factorial.Get_factorial(n, ch)
	fmt.Printf("%d! = %d\n",n,<-ch)
}