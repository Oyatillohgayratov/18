package corw

import (
	"log"
	"os"
)

func File2(s string,ch2 chan []byte){

	err := os.WriteFile("files/file2.txt", []byte(s), 0666)
	if err != nil {
		log.Fatal(err)
	}

	data2, err := os.ReadFile("files/file2.txt")
		if err != nil {
			log.Fatal(err)
		}
		ch2 <- data2
}