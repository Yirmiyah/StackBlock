package database

import "golang.org/x/crypto/bcrypt"

func cryptedKey(key string) string {
	hasKey, err := bcrypt.GenerateFromPassword([]byte(key), 12)
	if err != nil {
		panic(err)
	}
	return string(hasKey)
}

func AddCodeText(name string, codeCrypted string, key string) {

	_, err := OpenDB().Exec("INSERT INTO users (username,codeCrypted, key) VALUES (?, ?, ?)", name, codeCrypted, cryptedKey(key))
	if err != nil {
		panic(err)
	}
}
