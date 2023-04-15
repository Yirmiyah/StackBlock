package main

import (
	"StackBlock/database"
	"StackBlock/engine"
	"StackBlock/lib"
	"fmt"
	"net/http"
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
	ServerWS()

}

func ServerWS() {
	http.HandleFunc("/", WebSocketHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to start server: ", err)
	}
}
