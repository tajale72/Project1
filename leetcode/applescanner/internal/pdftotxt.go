package internal

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func PdftoTxt() {

	root := "/Users/romittajale/Documents/go/Project1/leetcode/applescanner/pdf"

	// Traverse the directory and process PDF files
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file is a regular file and has a PDF extension
		if !info.Mode().IsRegular() || filepath.Ext(path) != ".pdf" {
			return nil
		}

		// Read the PDF file
		inputBytes, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		outputFile := path + "output.txt"

		// Create a new file
		file, err := os.Create(outputFile)
		if err != nil {
			fmt.Println("Error:", err)

		}
		defer file.Close()

		scanner := bufio.NewScanner(strings.NewReader(string(inputBytes)))
		buf := make([]byte, 0, 1024*1024)
		scanner.Buffer(buf, 1024*1024)
		for scanner.Scan() {
			log.Println(scanner.Text())
			_, err = file.WriteString(scanner.Text())
			if err != nil {
				fmt.Println("Error:", err)
			}
		}

		return nil
	})

	if err != nil {
		log.Println("Error:", err)
	}
}
