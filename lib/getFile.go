package lib

import (
	"StackBlock/engine"
	"fmt"
	"os"
)

func GetFile() ([]byte, []byte) {
	argument := os.Args
	if len(argument) < 2 {
		panic("Please provide a file to read")

	}
	if len(argument) > 2 {
		panic("Please provide only one file to read")

	}

	// Open the file
	file, err := os.Open(argument[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// Read the file6
	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		panic(err)
	}
	// Print the file
	fmt.Println(string(buffer))
	return buffer, []byte(engine.GenerateRandomString())

}
