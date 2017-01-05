package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/bogem/id3v2"
)

const coverFilename string = "cover.jpg"

func addPicture() {
	filename := "Blure - Branches.mp3"

	frontCover, err := os.Open(coverFilename)
	if err != nil {
		log.Fatal("error opening cover file", err)
	}
	defer frontCover.Close()

	frontCoverBytes, err := ioutil.ReadAll(frontCover)
	if err != nil {
		log.Fatal("error reading cover file", err)
	}

	pic := id3v2.PictureFrame{
		Encoding:    id3v2.ENUTF8,
		MimeType:    "image/jpeg",
		Picture:     frontCoverBytes,
		PictureType: id3v2.PTFrontCover,
		Description: "Cover",
	}

	// Open file and find tag in it
	tag, err := id3v2.Open(filename)
	if err != nil {
		log.Fatal("error opening mp3 file", err)
	}
	defer tag.Close()

	tag.AddAttachedPicture(pic)
	if err = tag.Save(); err != nil {
		log.Fatal("Error while saving a tag:", err)
	}
}
