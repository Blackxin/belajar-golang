package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"simplified-fan-out-fan-in-pipeline/constants"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateDummyFilesSequentially() {
	log.Println("start")
	start := time.Now()

	generateFiles()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func generateFiles() {
	os.RemoveAll(constants.TempPath)
	os.MkdirAll(constants.TempPath, os.ModePerm)

	for i := 0; i < constants.TotalFile; i++ {
		fileName := filepath.Join(constants.TempPath, fmt.Sprintf("file-%d.txt", i))
		content := randomString(constants.ContentLength)
		err := os.WriteFile(fileName, []byte(content), os.ModePerm)
		if err != nil {
			log.Println("Error writing file", fileName)
		}

		log.Println(i, "files created")
	}

	log.Printf("%d of total files created", constants.TotalFile)
}
