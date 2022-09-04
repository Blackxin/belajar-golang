package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("random ke-1", rand.Int())
	fmt.Println("random ke-2", rand.Int())
	fmt.Println("random ke-3", rand.Int())

	// A.39.5. Angka Random Index Tertentu
	fmt.Println("random ke-4", rand.Intn(100))

	// A.39.6. Random Tipe Data String
	fmt.Println("random ke-5", randomString(5))
}

func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
