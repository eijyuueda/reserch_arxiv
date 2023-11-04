package pkg

import (
	"log"
	"os"
)

// テキストを書く
func WriteToFile(filename, text string) {
	// ファイルがなければ作る、あれば開く
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// テキストの書き込み
	_, err = file.WriteString(text + "\n")
	if err != nil {
		log.Fatal(err)
	}
}

// タイトルを書く
func WriteTitle(filename, text string) {
	// ファイルがなければ作る、あれば開く
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// テキストの書き込み
	_, err = file.WriteString("# " + text + "\n")
	if err != nil {
		log.Fatal(err)
	}
}

// サブタイトルを書く
func WriteSubTitle(filename, text string) {
	// ファイルがなければ作る、あれば開く
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// テキストの書き込み
	_, err = file.WriteString("## " + text + "\n")
	if err != nil {
		log.Fatal(err)
	}
}
