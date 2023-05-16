package configcrawler

import (
	"io"
	"net/http"

	// "golang.org/x/net/html"
)

type CSSFile struct {
	URL  string `json:"url"`
	Body string `json:"body"`
}

type PageData struct {
	HTMLContent string     `json:"html_content"`
	CSSFiles    []CSSFile  `json:"css_files"`
}

func downloadFile(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	fileContent, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(fileContent), nil
}

// func findAndDownloadCSS(url, htmlContent string) ([]CSSFile, error) {
// 	// TODO: Phân tích nội dung HTML và tìm các tệp CSS
// 	// Tải xuống từng tệp CSS và trả về một slice các CSSFile đã tải xuống

// 	var cssFiles []CSSFile

// 	tokenizer := html.NewTokenizer(response.Body)
// 	cssURL := "https://example.com/style.css"
// 	cssBody, err := downloadFile(cssURL)
// 	if err != nil {
// 		return nil, err
// 	}

// 	cssFile := CSSFile{
// 		URL:  cssURL,
// 		Body: cssBody,
// 	}
// 	cssFiles = append(cssFiles, cssFile)

// 	return cssFiles, nil
// }