package misollar

import (
	"fmt"
	"homework/corw"
)

func Masala2() {
	var s1,s2 string
	ch1 := make(chan []byte)
	ch2 := make(chan []byte)

	fmt.Print("file1ga malumot kiriting: ")
	fmt.Scan(&s1)
	fmt.Print("file2ga malumot kiriting: ")
	fmt.Scan(&s2)
	
	go corw.File1(s1,ch1)
	
	go corw.File2(s2,ch2)

	data := append(<-ch1, <-ch2...)

	corw.File3(data)

	fmt.Println("Ma'lumotlar yig'ib chiqarildi va file3.txt fayliga yozildi.")
}
