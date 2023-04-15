package database

import "golang.org/x/crypto/bcrypt"

func FindCodeText(name string) (string, string) {
	var codeCrypted string
	var key string
	err := OpenDB().QueryRow("SELECT codeCrypted, key FROM users WHERE username = ?", name).Scan(&codeCrypted, &key)
	if err != nil {
		panic(err)
	}

	//retrive the key from the database
	bcrypt.CompareHashAndPassword([]byte(key), []byte("password"))

}
