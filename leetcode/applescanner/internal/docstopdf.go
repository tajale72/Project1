package internal

import (
	"archive/zip"
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func DocsToPdf() {
	log.Printf("I am here in docs to pdf")
	// Open the docs file
	docPath := "/Users/romittajale/Documents/go/Project1/leetcode/applescanner/docs/apple.docx"
	path := "/Users/romittajale/Documents/go/Project1/leetcode/applescanner/docs/"
	doc, err := zip.OpenReader(docPath)
	if err != nil {
		log.Println("Error opening docs file:", err)
		return
	}
	defer doc.Close()

	// Read the contents of the docs file
	var text string
	for _, f := range doc.File {
		log.Println("f : ", f)
		if strings.HasSuffix(f.Name, ".xml") {
			rc, err := f.Open()
			if err != nil {
				log.Println("Error opening file in docs archive:", err)
				return
			}
			defer rc.Close()

			scanner := bufio.NewScanner(rc)
			for scanner.Scan() {
				log.Println("scanner.Text()", scanner.Text())
				text += scanner.Text()
			}
		}
	}

	// Create a new PDF document
	pdfPath := filepath.Join(filepath.Dir(path), "output.pdf")
	pdf, err := os.Create(pdfPath)
	if err != nil {
		log.Println("Error creating PDF file:", err)
		return
	}
	defer pdf.Close()

	// Write the text to the PDF document
	writer := bufio.NewWriter(pdf)
	writer.WriteString(text)
	writer.Flush()

	log.Println("PDF file saved to", pdfPath)
}
