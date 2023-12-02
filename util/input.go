package util

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	urlPattern       = "https://adventofcode.com/2023/day/%d/input"
	cacheFilePattern = "inputs/day%02d.input"
)

func GetPuzzleInput(day int) []string {
	cacheFile := fmt.Sprintf(cacheFilePattern, day)

	_, err := os.Stat(cacheFile)
	if os.IsNotExist(err) {
		fmt.Println("Input not cached, downloading...")
		input := getInputNetwork(day)
		defer input.Close()

		// cache the file
		f, err := os.Create(cacheFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		// stream the input to the cache file as we consume it
		return consumeLines(io.TeeReader(input, f))
	} else if err != nil {
		panic(err)
	} else {
		f, err := os.Open(cacheFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		return consumeLines(f)
	}
}

func getInputNetwork(day int) io.ReadCloser {
	session := os.Getenv("SESSION")
	if session == "" {
		panic("SESSION environment variable empty")
	}

	client := &http.Client{}

	url := fmt.Sprintf(urlPattern, day)
	req, _ := http.NewRequest("GET", url, nil)
	req.AddCookie(&http.Cookie{Name: "session", Value: session})
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return resp.Body
}

func consumeLines(r io.Reader) []string {
	scanner := bufio.NewScanner(r)

	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}
