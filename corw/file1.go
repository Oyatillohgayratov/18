package corw

import (
	"log"
	"os"
)

func File1(s string,ch1 chan []byte){

	err := os.WriteFile("files/file1.txt", []byte(s), 0666)
	if err != nil {
		log.Fatal(err)
	}

	data1, err := os.ReadFile("files/file1.txt")
		if err != nil {
			log.Fatal(err)
	}
		ch1 <- data1
}