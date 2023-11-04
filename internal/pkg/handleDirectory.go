package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

// ディレクトリの存在を確認し、存在しなければディレクトリを作成する。
func HandleDirectory(dirName string) {
	if _, err := os.Stat(dirName); err == nil {
		fmt.Println("Directory already exists:", dirName)
		return
	}

	fmt.Println("Making Directory:", dirName)
	err := os.Mkdir(dirName, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
}

// mdファイルの削除
func DeleteMarkdownFiles(dirName string) error {
	// ディレクトリ内のファイルを取得
	fileList, err := filepath.Glob(filepath.Join(dirName, "*.md"))
	if err != nil {
		return err
	}

	// 各ファイルを削除
	for _, file := range fileList {
		err := os.Remove(file)
		if err != nil {
			return err
		}
		fmt.Println("Deleted:", file)
	}
	return nil
}
