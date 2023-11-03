package pkg

import "strings"

func CleanFilename(filename string) string {
	invalidChars := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"} // 無効な文字リスト

	for _, c := range invalidChars {
		filename = strings.ReplaceAll(filename, c, "：") // 無効な文字を置換
	}
	return filename
}
