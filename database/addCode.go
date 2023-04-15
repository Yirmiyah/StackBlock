package database

func AddCode(codeCrypted string, key string) {

	_, err := OpenDB().Exec("INSERT INTO users (codeCrypted, key) VALUES (?, ?)", codeCrypted, key)
	if err != nil {
		panic(err)
	}
}
