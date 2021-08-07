package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	initRedis()
}

/*
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
*/

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})


	s := &http.Server{
		Addr:           ":8080",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

	var wordFlag string
	// var defFlag string

	if len(os.Args) >= 1 {
		wordFlag = os.Args[0]
		// defFlag = os.Args[2]
		flag.Parse()
	} else {
		log.Fatal("Please specify a word and definition first.")
	}

	ok, err := Exists(wordFlag)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(ok)
}
