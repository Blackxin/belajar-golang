package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	totalFile       = 3000
	contentLength   = 5000
	timeOutDuration = 15 * time.Second
)

var tempPath = filepath.Join(os.Getenv("XDG_CACHE_HOME"), "chapter-A.61-pipeline-cancellation-context")

type FileInfo struct {
	Index       int
	FileName    string
	WorkerIndex int
	Err         error
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// main
func generateDummyFilesConcurrently() {
	log.Println("start")
	start := time.Now()

	// generateFiles()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	time.AfterFunc(timeOutDuration, cancel)
	generateFilesWithContext(ctx)

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
	generateFilesWithContext(context.Background())
}

func generateFilesWithContext(ctx context.Context) {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	done := make(chan int)

	go func() {
		chanFileIndex := generateFileIndexes(ctx)

		createFilesWorker := 100
		chanFileResult := createFiles(ctx, chanFileIndex, createFilesWorker)

		counterSuccess := 0
		for fileResult := range chanFileResult {
			if fileResult.Err != nil {
				log.Printf("error creating file %s. stack trace: %s", fileResult.FileName, fileResult.Err)
			} else {
				counterSuccess++
			}
		}
	}()

	select {
	case <-ctx.Done():
		log.Printf("generation process stopped %s", ctx.Err())
	case counterSuccess := <-done:
		log.Printf("%d/%d of total files created", counterSuccess, totalFile)
	}

}

func generateFileIndexes(ctx context.Context) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	go func() {
		for i := 0; i < totalFile; i++ {
			select {
			case <-ctx.Done():
				break
			default:
				chanOut <- FileInfo{
					Index:    i,
					FileName: fmt.Sprintf("file-%d.txt", i),
				}
			}

		}
		close(chanOut)
	}()

	return chanOut
}

func createFiles(ctx context.Context, chanIn <-chan FileInfo, numberOfWorkers int) <-chan FileInfo {
	chanOut := make(chan FileInfo)

	wg := new(sync.WaitGroup)
	wg.Add(numberOfWorkers)

	go func() {
		for workerIndex := 0; workerIndex < numberOfWorkers; workerIndex++ {
			go func(workerIndex int) {
				for job := range chanIn {
					select {
					case <-ctx.Done():
						break
					default:
						filePath := filepath.Join(tempPath, job.FileName)
						content := randomString(contentLength)
						err := os.WriteFile(filePath, []byte(content), os.ModePerm)

						log.Println("worker", workerIndex, "working on", job.FileName, "file generation")

						chanOut <- FileInfo{
							FileName:    job.FileName,
							WorkerIndex: workerIndex,
							Err:         err,
						}
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
