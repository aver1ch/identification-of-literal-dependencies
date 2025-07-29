package main

import (
	"fmt"
	ap "identification-of-literal-dependencies/internal/antiPlagiat"
	readFiles "identification-of-literal-dependencies/internal/readTxtFiles"
)

func main() {
	docs, err := readFiles.GetTxtFiles("./data/converted_txt")
	if err != nil {
		fmt.Errorf("Error: %w", err)
	}

	n, w := 5, 4

	for _, doc := range docs {
		_ = ap.IndexDocument(doc, n, w)
	}

	newText, _ := ap.ReadAndClean("./data/input/input.txt")
	matches := ap.CheckPlagiarism(newText, n, w)
	ap.ReportMatches(matches)
}
