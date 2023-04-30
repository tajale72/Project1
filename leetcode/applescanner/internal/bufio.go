package internal

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func BufioScanningDevice() (int, int) {
	path := "/Users/romittajale/Documents"
	password := "romit"
	regex := regexp.MustCompile(password)
	var numoffiles, numoffilesscanned int

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		fileContents, err := ioutil.ReadFile(path)
		if err != nil {
			log.Printf("error from reading the files %s", err)
			return err
		}
		numoffilesscanned++
		lineNumber := 0
		scanner := bufio.NewScanner(strings.NewReader(string(fileContents)))
		buf := make([]byte, 0, 1024*1024)
		scanner.Buffer(buf, 1024*1024)
		for scanner.Scan() {
			if regex.MatchString(scanner.Text()) {
				log.Printf("Found password '%s' in file '%s' at line %d", password, path, lineNumber)
				numoffiles++
			}
		}

		if err := scanner.Err(); err != nil {
			log.Printf("error from scanner %s", err)
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
	return numoffiles, numoffilesscanned
}
