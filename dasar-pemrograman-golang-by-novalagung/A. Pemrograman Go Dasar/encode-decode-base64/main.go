package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// A.46.1. Penerapan Fungsi EncodeToString() & DecodeString()
	// data := "taufik hidayat"
	// encodedString := base64.StdEncoding.EncodeToString([]byte(data))
	// fmt.Println("encoded:", encodedString)

	// decodedByte, err := base64.StdEncoding.DecodeString(encodedString)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// decodedString := string(decodedByte)
	// fmt.Println("decoded:", decodedString)

	// A.46.2. Penerapan Fungsi Encode() & Decode()
	// encoded := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	// base64.StdEncoding.Encode(encoded, []byte(data))
	// encodedString := string(encoded)
	// fmt.Println(encodedString)

	// decoded := make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	// _, err := base64.RawURLEncoding.Decode(decoded, encoded)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// decodedString := string(decoded)
	// fmt.Println(decodedString)

	// A.46.3. Encode & Decode Data URL
	data := "https://tfkhdyt.my.id"

	encodedString := base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedString)

	decodedByte, err := base64.URLEncoding.DecodeString(encodedString)
	if err != nil {
		panic(err.Error())
	}
	decodedString := string(decodedByte)
	fmt.Println(decodedString)
}
