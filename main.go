package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
	"sync"
)

const (
	settingsDownloadsKey = "DOWNLOADSDIR"
	settingsSessionId    = "SESSIONID"
)

var (
	settings        map[string]string
	scanner         *bufio.Scanner
	finalizeWithZip bool
	configPath      string
)

func main() {
	// settings defaults
	settings = map[string]string{
		settingsDownloadsKey: "downloads",
		settingsSessionId:    "",
	}
	scanner = bufio.NewScanner(os.Stdin)

	// get application dir
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	configPath = usr.HomeDir + "/.SGo-Scraper"

	if _, err := os.Stat(configPath); err == nil {
		settings, err = godotenv.Read(configPath)
		if err != nil {
			panic(err)
		}
	}

	finalizeWithZip = *flag.Bool("zip", false, "zip")
	flag.Parse()

	for {
		getUrl()
	}
}

func getUrl() {
	fmt.Print("Please enter album URL: ")
	scanner.Scan()
	albumURL := scanner.Text()

	pageSource := getContents(albumURL)
	modelName, albumName := getAlbumInfo(pageSource)
	imagesFound := crawlImages(pageSource)

	for len(imagesFound) <= 0 {
		fmt.Println("No images found.  Session may be expired, enter new SessionID: ")
		scanner.Scan()
		settings[settingsSessionId] = scanner.Text()
		pageSource = getContents(albumURL)
		modelName, albumName = getAlbumInfo(pageSource)
		imagesFound = crawlImages(pageSource)
	}

	fmt.Println("Found", albumName, "set from", modelName, "!")
	fmt.Println("Found", len(imagesFound), "images in set. Downloading...")

	albumDir := settings[settingsDownloadsKey] + "/" + modelName + " - " + albumName

	checkAndCreateDir(settings[settingsDownloadsKey])
	checkAndCreateDir(albumDir)
	imagesDownloaded := make([]string, 0)

	godotenv.Write(settings, configPath)

	fmt.Print(strings.Repeat(".", len(imagesFound)))
	fmt.Print("\r")

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
	b, _ := saveImage(imageURL, outputUrl)
	if b > 0 {
		fmt.Print("✓")
	} else {
		fmt.Print("✗")
	}
}
