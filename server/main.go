package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/wzslr321/words/cli"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

func init() {
	// initRedis()
}

var dir = "/server/words/"

func main() {

	actions := cli.RegisterActions()
	actions.Add = os.Args[1]
	var word string

	if len(os.Args) > 1 {
		word = os.Args[1]
		flag.Parse()
	}

	if word != "" {
		if err := addNewWords(word); err != nil {
			fmt.Printf("Error occured while adding new words: %v", err)
		}
	} else {
		// router := initRouter()

		s := &http.Server{
			Addr: ":8080",
			// Handler:        router,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}

		s.ListenAndServe()
	}
}

func addNewWords(word string) error {
	_, b, _, _ := runtime.Caller(0)
	basePath := filepath.Dir(b)
	localDir := filepath.Dir(basePath)

	path := localDir + "/server/words/new_words.txt"

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	defer file.Close()
	if err != nil {
		log.Printf("Unable to open words.txt: %v", err)
	}

	scanner := bufio.NewScanner(file)
	var doesExist bool
	for scanner.Scan() {
		line := scanner.Text()
		if bytes.Contains([]byte(line), []byte(word)) {
			doesExist = true
		}
	}
	if err = scanner.Err(); err != nil {
		log.Printf("Scanner returned error: %v", err)
	}
	if doesExist {
		log.Fatalf("This word already exists in words.txt file")
	}

	writer := bufio.NewWriter(file)

	if _, err = writer.WriteString(word + "\n"); err != nil {
		log.Printf("Unable to write new word. \n Error: %v", err)
	}

	err = writer.Flush()
	if err != nil {
		log.Printf("Failed to execute .Flush() method.\n Error: %v", err)
	}

	return err
}
