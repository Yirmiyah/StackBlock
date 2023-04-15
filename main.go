package main

import (
	"StackBlock/database"
	"StackBlock/engine"
	"StackBlock/lib"
	"fmt"
)

func main() {
	database.OpenDB()
	defer database.OpenDB().Close()

	text, textKey := lib.GetFile()
	fmt.Printf("Text: %v\n", string(text))

	encryptText, err := engine.EncryptText(text, textKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Encrypted text: %v\n", encryptText)

	decryptedText, err := engine.DecryptText(encryptText, textKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted text: %v\n", decryptedText)

}
