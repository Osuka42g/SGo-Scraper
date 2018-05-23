package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	downloadsDir := os.Getenv("DOWNLOADSDIR")
	args := os.Args
	albumURL := args[1]
	finalizeWithZip := args[len(args)-1] == "-z"

	pageSource := getContents(albumURL)
	modelName, albumName := getAlbumInfo(pageSource)
	imagesFound := crawlImages(pageSource)

	fmt.Println("Found", albumName, "set from", modelName, "!")
	fmt.Println("Found", len(imagesFound), "images in set. Downloading...")

	albumDir := downloadsDir + "/" + modelName + " - " + albumName

	checkAndCreateDir(downloadsDir)
	checkAndCreateDir(albumDir)
	imagesDownloaded := make([]string, 0)

	var wg sync.WaitGroup
	wg.Add(len(imagesFound))

	countSize := digitsLen(len(imagesFound))
	for i, imageURL := range imagesFound {
		imageOutput := albumDir + "/" + leftPad(strconv.Itoa(i+1), "0", countSize) + ".jpg"
		go getFile(&wg, imageURL, imageOutput)
		imagesDownloaded = append(imagesDownloaded, imageOutput)
	}

	wg.Wait()

	if finalizeWithZip {
		err := ZipFiles(albumDir+"/"+albumName+".zip", imagesDownloaded)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("")
	fmt.Println("Done... Enjoy!")
	fmt.Println(albumDir)
}

func getFile(wg *sync.WaitGroup, imageURL string, outputUrl string) {
	defer wg.Done()
	fmt.Print(".")
	b, _ := saveImage(imageURL, outputUrl)
	if b > 0 {
		fmt.Print("✓")
	} else {
		fmt.Print("✗")
	}
}
