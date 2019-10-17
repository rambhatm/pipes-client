package main

import (
	"log"
	"os"
)

type pipeMetadata struct {
	path string
}

func fileInfo(path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	statData, err := file.Stat()
	log.Printf("%#v", statData)

	return
}

func main() {
	fileInfo("./client.go")
}
