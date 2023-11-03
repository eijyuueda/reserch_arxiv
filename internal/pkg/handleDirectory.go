package pkg

import (
	"fmt"
	"os"
)

func HandleDirectory(dirName string) {
	// ディレクトリの存在を確認し、存在する場合は削除する
	if _, err := os.Stat(dirName); err == nil {
		err := os.RemoveAll(dirName)
		if err != nil {
			fmt.Println("Error deleting existing directory:", err)
			return
		}
	}

	// 新しいディレクトリを作成する
	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
}
