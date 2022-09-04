package main

import (
	"crypto/md5"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

var _tempPath = filepath.Join(os.Getenv("XDG_CACHE_HOME"), "chapter-A.59-pipeline-temp")

func __main() {
	log.Println("start")
	start := time.Now()

	proceed()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func proceed() {
	counterTotal := 0
	counterRenamed := 0
	err := filepath.Walk(_tempPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		counterTotal++

		buf, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		sum := fmt.Sprintf("%x", md5.Sum(buf))

		destinationPath := filepath.Join(_tempPath, fmt.Sprintf("file-%x.txt", sum))
		err = os.Rename(path, destinationPath)
		if err != nil {
			return err
		}

		counterRenamed++
		return nil
	})

	if err != nil {
		log.Println("ERROR:", err.Error())
	}

	log.Printf("%d/%d files renamed", counterRenamed, counterTotal)
}
