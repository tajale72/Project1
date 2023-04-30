package internal

import (
	"log"
	"os"
	"strings"
)

func GetRootPath() string {
	env := os.Environ()
	var val []string
	for _, v := range env {
		val = strings.Split(v, "=")
		if val[0] == "HOME" {
			newStr := strings.ReplaceAll(val[1], "/Users/", "")
			log.Printf("Hi %s", newStr)
		}
	}
	return val[1]
}
