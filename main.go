package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	wordFlag := flag.String("word", "", "word name")
	defFlag := flag.String("def", "", "definition of the word")

	flag.Parse()

	if *wordFlag == "" || *defFlag == "" {
		log.Fatal("Please specify a word and definition first.")
	}

	file, err := os.Open("words.txt")
	if err != nil {
		log.Printf("Unable to open words.txt: %v", err)
	}

	scanner := bufio.NewScanner(file)

	var doesExist bool

	for scanner.Scan() {
		line := scanner.Text()
		if bytes.Contains([]byte(line), []byte(*wordFlag)) {
			doesExist = true
		}
	}

	if err = scanner.Err(); err != nil {
		log.Printf("Scanner returned error: %v", err)
	}

	file.Close()

	if doesExist {
		log.Fatalf("This word already exists in words.txt file")
	}

	newWord := *wordFlag + " - " + *defFlag + "\n"

	file, err = os.OpenFile("words.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error creating a file: %v", err)
	}
	writer := bufio.NewWriter(file)

	writer.WriteString(newWord)

	writer.Flush()
	file.Close()
}
