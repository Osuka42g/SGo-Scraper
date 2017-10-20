package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	albumURL := os.Args[1]

	ch := make(chan []string)

	go crawlImages(albumURL, ch)
	imagesFound := <-ch

	fmt.Println("\nFound", len(imagesFound), "images in set. Downloading...")

	checkAndCreateDir("downloads")
	checkAndCreateDir("downloads/model - set")

	for i, imageURL := range imagesFound {
		fmt.Println(" - " + imageURL)

		imageOutput := "downloads/model - set/" + leftPad(strconv.Itoa(i), "0", digitsLen(i)%2) + ".jpg"
		fmt.Println(imageOutput)
		img, _ := os.Create(imageOutput)
		resp, _ := http.Get(imageURL)

		b, _ := io.Copy(img, resp.Body)
		fmt.Println("File size:", b)
	}

	fmt.Println("Done... Enjoy!")
}
