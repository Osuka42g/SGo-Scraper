package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	return strings.Repeat(padStr, pLen) + s
}

func saveImage(url string, output string) (int64, error) {
	img, _ := os.Create(output)
	resp, _ := http.Get(url)
	return io.Copy(img, resp.Body)
}
