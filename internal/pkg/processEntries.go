package pkg

import (
	"path/filepath"
	"strings"
	"time"
)

func ProcessEntries(feed Feed) {
	today := time.Now().Format("20060102")
	// 概要ファイルの英語、日本語版
	english_ab := "abstract.md"
	japanese_ab := "abstract_ja.md"
	for _, entry := range feed.Entries {
		// タイトルから改行を取り除く
		title := strings.ReplaceAll(entry.Title, "\n", "")
		// 原文のリンク集
		WriteToFile(filepath.Join("docs", today, english_ab), title)
		WriteToFile(filepath.Join("docs", today, english_ab), entry.ID)
		// 翻訳後の概要とタイトル
		translateText, _ := Translate(entry.Summary, "en", "ja")
		translateText = strings.ReplaceAll(translateText, "\n", "")
		WriteToFile(filepath.Join("docs", today, japanese_ab), title)
		WriteToFile(filepath.Join("docs", today, japanese_ab), entry.ID)
		WriteToFile(filepath.Join("docs", today, japanese_ab), translateText)
		// DLの前処理
		processEntry(entry, today, title)
	}
}

// エントリーのhrefを取り出す処理
func processEntry(entry Entry, today, title string) {
	for _, link := range entry.Links {

		if link.Rel == "related" {
			processRelatedLink(title, link.HRef, today)
		}
	}
}

// hrefからDLをする
func processRelatedLink(title, href, today string) {
	filename := CleanFilename(title + ".pdf")
	// fmt.Println(title)
	DownloadPDF(href, filepath.Join("docs", today, filename))
}
