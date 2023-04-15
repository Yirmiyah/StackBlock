package database

func AddCodeText(name string, codeCrypted string, key string) {

	_, err := OpenDB().Exec("INSERT INTO users (username,codeCrypted, key) VALUES (?, ?, ?)", name, codeCrypted, key)
	if err != nil {
		panic(err)
	}
}
