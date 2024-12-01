package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	t := time.Now()
	day := t.Day()
	dayStr := fmt.Sprintf("%02d", day)
	inputFilename := fmt.Sprintf("%s/input.txt", dayStr)
	testFilename := fmt.Sprintf("%s/test.txt", dayStr)
	scriptFilename := fmt.Sprintf("%s/main.go", dayStr)
	templateFilename := "template.go.txt"

	if _, err := os.Stat(dayStr); err == nil {
		fmt.Printf("Directory '%s' already exists, exiting.\n", dayStr)
		return
	}

	mkdirp(dayStr)

	inputFile, err := os.Create(inputFilename)
	fail(err)
	defer inputFile.Close()

	testFile, err := os.Create(testFilename)
	fail(err)
	defer testFile.Close()

	scriptFile, err := os.Create(scriptFilename)
	fail(err)
	defer scriptFile.Close()
	templateFile, err := os.Open(templateFilename)
	fail(err)

	if _, err := io.Copy(scriptFile, templateFile); err != nil {
		fail(err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", t.Day()), nil)
	fail(err)
	req.Header.Set("Cookie", "session="+os.Getenv("SESSION_COOKIE"))

	resp, err := http.DefaultClient.Do(req)
	fail(err)
	defer resp.Body.Close()

	if _, err := io.Copy(inputFile, resp.Body); err != nil {
		fail(err)
	}
}

func fail(err error) {
	if err != nil {
		panic(err)
	}
}

func mkdirp(dir string) {
	err := os.Mkdir(dir, 0755)
	if err != nil && !os.IsExist(err) {
		fail(err)
	}
}
