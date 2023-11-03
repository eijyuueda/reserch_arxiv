package main

import (
	"arxiv_search/internal/pkg"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		panic("No arguments provided. Program requires an argument(string).")
	}
	endpoint := "http://export.arxiv.org/api/query"
	query := "search_query=abs:" + os.Args[1]
	query = query + "&sortBy=submittedDate&sortOrder=descending&max_results=20"

	resp, err := http.Get(fmt.Sprintf("%s?%s", endpoint, query))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// For XML Response
	feed, err := pkg.ParseXML(body)
	if err != nil {
		fmt.Println("Error parsing XML:", err)
		return
	}

	// ディレクトリの作成
	today := time.Now().Format("20060102")
	pkg.HandleDirectory("docs")
	pkg.HandleDirectory(filepath.Join("docs", today))
	// エントリーの処理
	pkg.ProcessEntries(feed)
}
