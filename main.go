package main

import (
	"bufio"
	"bytes"
	"flag"
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
	"path/filepath"
)

type application struct {
	pool *redis.Pool
}

var app *application

func init() {
	app.initRedis()
}

func main() {

    var wordFlag string 
    var defFlag string
    
    if len(os.Args) >= 3 {
	    wordFlag = os.Args[1]
	    defFlag = os.Args[2]
	    flag.Parse()
    } else {
		log.Fatal("Please specify a word and definition first.")
    }


    var err error
    var dir string

	dir, err = filepath.Abs("development/words/words.txt")
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.OpenFile(dir, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Unable to open words.txt: %v", err)
	}

	scanner := bufio.NewScanner(file)

	var doesExist bool

	for scanner.Scan() {
		line := scanner.Text()
		if bytes.Contains([]byte(line), []byte(wordFlag)) {
			doesExist = true
		}
	}

	if err = scanner.Err(); err != nil {
		log.Printf("Scanner returned error: %v", err)
	}


	if doesExist {
		log.Fatalf("This word already exists in words.txt file")
	}

	newWord := wordFlag + " - " + defFlag + "\n"

	writer := bufio.NewWriter(file)

	if _, err = writer.WriteString(newWord); err != nil {
		log.Printf("Unable to write new word. \n Error: %v", err)
	}

	err = writer.Flush()
	if err != nil {
		log.Printf("Failed to execute .Flush() method.\n Error: %v", err)
	}

	_ = file.Close()
}
