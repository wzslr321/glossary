package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func initRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	r.POST("/:word", func(c *gin.Context) {
		word := c.Param("word")

		res, _ := Exists(word)

		if res == false {
			val := scrap(word)
			byteValue, _ := json.Marshal(val)
			Set(word, byteValue)

			file, err := os.OpenFile("words.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			defer file.Close()
			if err != nil {
				log.Printf("Unable to open words.txt: %v", err)
			}

			writer := bufio.NewWriter(file)

			if _, err = writer.WriteString(word + "\n"); err != nil {
				log.Printf("Unable to write new word. \n Error: %v", err)
			}

			err = writer.Flush()
			if err != nil {
				log.Printf("Failed to execute .Flush() method.\n Error: %v", err)
			}

		}
	})

	r.GET("/:word", func(c *gin.Context) {
		word := c.Param("word")
		res, _ := Get(word)

		var data []string

		if err := json.Unmarshal(res, &data); err != nil {
			log.Printf("Error occured while unmarshalling: %v", err)
		}

		c.JSON(http.StatusOK, gin.H{
			"definition": data,
		})
	})

	r.POST("/fetch", func(c *gin.Context) {
		if err := fetchNewWords(); err != nil {
			log.Printf("Unable to fetch words: %v", err)
		}
	})
	return r
}

func fetchNewWords() error {

	file, err := os.OpenFile("new_words.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	defer file.Close()
	if err != nil {
		log.Printf("Unable to open words.txt: %v", err)
	}

	scanner := bufio.NewScanner(file)
	var newWords []string

	for scanner.Scan() {
		key := scanner.Text()
		newWords = append(newWords, key)
		exists, _ := Exists(key)
		if !exists {
			val := scrap(key)
			byteValue, _ := json.Marshal(val)
			Set(key, byteValue)
		}
	}
	if err = scanner.Err(); err != nil {
		log.Printf("Scanner returned error: %v", err)
	}

	fetchNewWordsWithTxt(newWords)

	return err
}

func fetchNewWordsWithTxt(words []string) error {
	file, err := os.OpenFile("words.txt", os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		log.Printf("Unable to open words.txt: %v", err)
	}

	writer := bufio.NewWriter(file)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		for i, word := range words {
			if bytes.Contains([]byte(line), []byte(word)) {
				words[i] = words[len(words)-1]
				words = words[:len(words)-1]
			}
		}
	}
	if err = scanner.Err(); err != nil {
		log.Printf("Scanner returned error: %v", err)
	}

	for _, word := range words {
		if _, err = writer.WriteString(word + "\n"); err != nil {
			log.Printf("Unable to write new word. \n Error: %v", err)
		}
	}

	err = writer.Flush()
	if err != nil {
		log.Printf("Failed to execute .Flush() method.\n Error: %v", err)
	}

	return err
}
