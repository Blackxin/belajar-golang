package main

import (
	"crypto/sha1"
	"fmt"
	"time"
)

func main() {
	// A.47.1. Penerapan Hash SHA1
	text := "Nama saya Taufik Hidayat"
	// sha := sha1.New()
	// sha.Write([]byte(text))
	// encrypted := sha.Sum(nil)
	// encryptedString := fmt.Sprintf("%x", encrypted)

	// fmt.Println(encryptedString)

	// A.47.2. Metode Salting Pada Hash SHA1
	fmt.Printf("original : %s\n\n", text)

	hashed1, _ := doHashUsingSalt(text)
	fmt.Printf("hashed 1 : %s\n\n", hashed1)

	hashed2, _ := doHashUsingSalt(text)
	fmt.Printf("hashed 1 : %s\n\n", hashed2)

	hashed3, _ := doHashUsingSalt(text)
	fmt.Printf("hashed 1 : %s\n\n", hashed3)
}

func doHashUsingSalt(text string) (string, string) {
	salt := fmt.Sprintf("%d", time.Now().UnixNano())
	saltedText := fmt.Sprintf("text: '%s', salt: '%s'", text, salt)
	fmt.Println(saltedText)

	sha := sha1.New()
	sha.Write([]byte(saltedText))
	encrypted := sha.Sum(nil)

	return fmt.Sprintf("%x", encrypted), salt
}
