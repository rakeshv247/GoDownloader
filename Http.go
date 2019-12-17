package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const mp3File string = "song.mp3"

// GetMetaData makes a HEAD request and obtains the metadata of the requested url
func GetMetaData(url *string) int64 {
	fmt.Printf("Getting metadata of url %s\n", *url)

	resp, err := http.Head(*url)
	if err != nil {
		log.Fatalf("Failed receiving metadata of url %s\n", *url)
	}
	defer resp.Body.Close()

	return resp.ContentLength
}

// Download does the actual Download of the requested file
func Download(url *string) bool {
	fmt.Printf("In http GET func, received url: %s\n", *url)

	ContentLength := GetMetaData(url)
	if ContentLength == 0 {
		return false
	}

	resp, err := http.Get(*url)
	if err != nil {
		log.Fatalf("Failed on http GET for url %s\n", *url)
	}
	defer resp.Body.Close()

	file, err := os.Create(mp3File)
	if err != nil {
		log.Fatalf("Error opening file %s\n", mp3File)
	}
	defer file.Close()

	n, err := io.Copy(file, resp.Body)
	if err != nil || n != ContentLength {
		log.Fatalf("Error copying file %s\n", mp3File)
	}

	fmt.Println("File downloaded successfully...!!!")
	return true
}
