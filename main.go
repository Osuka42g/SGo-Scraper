package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	downloadsDir := "downloads"

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	albumURL := os.Args[1]

	pageSource := getContents(albumURL)
	imagesFound := crawlImages(pageSource)

	fmt.Println("Found", len(imagesFound), "images in set. Downloading...")

	modelName := getModelName(pageSource)
	albumName := getAlbumName(pageSource)

	albumDir := downloadsDir + "/" + modelName + " - " + albumName

	checkAndCreateDir(downloadsDir)
	checkAndCreateDir(albumDir)

	for i, imageURL := range imagesFound {
		imageOutput := albumDir + "/" + leftPad(strconv.Itoa(i), "0", digitsLen(len(imagesFound))-1) + ".jpg"
		fmt.Println(imageURL + " -> " + imageOutput)

		b, _ := saveImage(imageURL, imageOutput)
		fmt.Println("File size:", b)
	}

	fmt.Println("Done... Enjoy!")
}
