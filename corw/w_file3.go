package corw

import (
	"log"
	"os"
)

func File3(data []byte){
	err := os.WriteFile("files/file3.txt", data, 0666)
	if err != nil {
		log.Fatal(err)
	}
}