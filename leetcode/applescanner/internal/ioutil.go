package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func ScanningDevice() (int, int) {

	//path := GetRootPath()
	var numoffiles, numoffilesscanned int
	path := "/Users/romittajale/Documents"
	pass := "romit"
	regex := regexp.MustCompile(pass)

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}

		fileContents, err := ioutil.ReadFile(path)
		numoffilesscanned++
		if err != nil {
			log.Printf("error from reading the files %s", err)
			return err
		}

		lines := strings.Split(string(fileContents), "\n")
		for i, line := range lines {
			if regex.MatchString(line) {
				fmt.Printf("Found password '%s' in file '%s' on line %d\n", pass, path, i+1)
				numoffiles++
			}
		}

		return nil

	})
	if err != nil {
		log.Printf("error from file path %s", err)
	}
	return numoffiles, numoffilesscanned
}
