package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"applescanner/internal"
)

func main() {

	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	r.GET("/ioutil", ReadAllScan)
	r.GET("/bufio", BufioScan)
	r.GET("/pdftotxt", PdftoTxt)
	r.POST("/upload", DocstoPdf)
	r.Run()

}

type Scanner struct {
	Str                  string `json:"numberoffiles"`
	NumberoffilesScanned string `json:"numberoffilesScanned"`
	Elapsed              time.Duration
}

func ReadAllScan(c *gin.Context) {
	var scanner Scanner
	start := time.Now() // Start the timer before calling the function          // Call the function
	numOffiles, numoffilesscanned := internal.ScanningDevice()
	scanner.Str = fmt.Sprintf("Password found in %d files", numOffiles)

	Elapsed := time.Since(start) // Calculate the time Elapsed since the timer was started
	scanner.Elapsed = time.Duration(Elapsed.Microseconds())

	scanner.NumberoffilesScanned = fmt.Sprintf("Total number of files scanned %d", numoffilesscanned)
	log.Printf("Time Elapsed: %v\n", scanner)
	c.JSON(http.StatusAccepted, scanner)
}

func BufioScan(c *gin.Context) {
	var scanner Scanner
	start := time.Now() // Start the timer before calling the function          // Call the function
	numOffiles, numoffilesscanned := internal.BufioScanningDevice()
	scanner.Str = fmt.Sprintf("Password found in %d files", numOffiles)

	Elapsed := time.Since(start) // Calculate the time Elapsed since the timer was started
	scanner.Elapsed = time.Duration(Elapsed.Microseconds())

	scanner.NumberoffilesScanned = fmt.Sprintf("Total number of files scanned %d", numoffilesscanned)

	log.Printf("Time Elapsed: %v\n", scanner)
	c.JSON(http.StatusAccepted, scanner)
}

func PdftoTxt(c *gin.Context) {
	internal.PdftoTxt()

}

// DocstoPDf is a function is
func DocstoPdf(c *gin.Context) {
	log.Println("uploading a file.....")
	if c.Request.Method != http.MethodPost {
		c.JSON(http.StatusMethodNotAllowed, "Not allowed")
		return
	}

	//Getting the input file from the request
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	//Getting the filename
	filename := filepath.Base(header.Filename)
	//Creating a file and uploading the file inside the uploads folder
	out, err := os.Create(filepath.Join("uploads", filename))
	if err != nil {
		log.Println("serror from creating the file", err)
	}
	defer out.Close()

	// Copy contents from uploaded file to new file on disk
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	// Convert the PDF to DOCX
	successmsg := "File uploaded successfully"
	// Return success response
	log.Println(successmsg)
	c.JSON(http.StatusOK, gin.H{
		"message": successmsg,
	})
	//internal.DocsToPdf()
}
