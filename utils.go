package main

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"strconv"
	"fmt"
)

func checkAndCreateDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, os.ModePerm)
	}
}

func digitsLen(n int) int {
	return len(strconv.Itoa(n))
}

func leftPad(s string, padStr string, pLen int) string {
	return fmt.Sprintf("%"+padStr+strconv.Itoa(pLen)+"s", s)
}

func saveImage(url string, output string) (int64, error) {
	img, err := os.Create(output)
	if err != nil {
		return 0, err
	}
	defer img.Close()
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return io.Copy(img, resp.Body)
}

func ZipFiles(filename string, files []string) error {

	newfile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newfile.Close()

	zipWriter := zip.NewWriter(newfile)
	defer zipWriter.Close()

	for _, file := range files {

		zipfile, err := os.Open(file)
		if err != nil {
			return err
		}
		//noinspection GoDeferInLoop
		defer zipfile.Close()

		info, err := zipfile.Stat()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, zipfile)
		if err != nil {
			return err
		}
	}
	return nil
}
