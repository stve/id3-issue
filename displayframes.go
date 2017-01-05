package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/bogem/id3v2"
)

func main() {
	// addPicture()

	fmt.Println("Detecting MP3 files...")
	files, _ := filepath.Glob("*.mp3")
	for i := 0; i < len(files); i++ {
		fmt.Println("********************")
		fmt.Println("File: " + files[i] + ":")

		// Open file and find tag in it
		tag, err := id3v2.Open(files[i])
		if err != nil {
			log.Fatal("error opening file", err)
		}

		defer tag.Close()

		frameList := tag.AllFrames()
		for key, value := range frameList {
			fmt.Println(key)
			fmt.Println(value)
		}
	}
}
