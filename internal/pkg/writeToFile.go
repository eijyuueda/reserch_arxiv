package pkg

import (
	"log"
	"os"
)

func WriteToFile(filename, translation string) {
	// ファイルがなければ作る、あれば開く
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(translation + "\n")
	if err != nil {
		log.Fatal(err)
	}
}
