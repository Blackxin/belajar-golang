package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"simplified-fan-out-fan-in-pipeline/constants"
)

type FileInfo struct {
	Index       int
	FileName    string
	WorkerIndex int
	Err         error
}

func generateDummyFilesConcurrently() {
	log.Println("start")
	start := time.Now()

	generateFilesConcurrently()

	duration := time.Since(start)
	log.Println("done in", duration.Seconds(), "seconds")
}

func generateFilesConcurrently() {
	os.RemoveAll(constants.TempPath)
	os.MkdirAll(constants.TempPath, os.ModePerm)

	// pipeline 1: job distribution
	chanFileIndex := generateFileIndexes()

	// pipeline 2: the main logic (creating files)
	createFileWorker := 100
	chanFileResult := createFiles(chanFileIndex, createFileWorker)

	// track and print output
	counterTotal := 0
	counterSuccess := 0

	for fileResult := range chanFileResult {
		if fileResult.Err != nil {
			log.Printf("error creating file %s, stack trace: %s", fileResult.FileName, fileResult.Err)
		} else {
			counterSuccess++
		}
		counterTotal++
	}

	log.Printf("%d/%d of total files created", counterSuccess, counterTotal)
}

func generateFileIndexes() <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for i := 0; i < constants.TotalFile; i++ {
			chanOut <- FileInfo{
				Index:    i,
				FileName: fmt.Sprintf("file-%d.txt", i),
			}
		}
		close(chanOut)
	}()

	return chanOut
}

func createFiles(chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	// wait group to control the workers
	wg := new(sync.WaitGroup)

	// allocate N of workers
	wg.Add(numberOfWorkers)

	go func() {
		// dispatch N workers
		for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
			go func(workerIndex int) {
				// listen to chanIn for incoming jobs
				for job := range chanIn {
					// do the jobs
					filePath := filepath.Join(constants.TempPath, job.FileName)
					content := randomString(constants.ContentLength)
					err := os.WriteFile(filePath, []byte(content), os.ModePerm)

					log.Println("worker", workerIndex, "working on", job.FileName, "file generation")

					// construct the job's result, and send it to 'chanOut'
					chanOut <- FileInfo{
						FileName:    job.FileName,
						WorkerIndex: workerIndex,
						Err:         err,
					}
				}

				wg.Done()
			}(workerIndex)
		}
	}()

	go func() {
		wg.Wait()
		close(chanOut)
	}()

	return chanOut
}
