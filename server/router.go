package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func initRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())

	r.POST("/:word", func(c *gin.Context) {
		word := c.Param("word")

		res, err := Get(word)
		if err != nil {
			log.Println(err)
		}

		if len(res) == 0 {
			val := scrap(word)
			byteValue, _ := json.Marshal(val)
			Set(word, byteValue)
		}
	})

	r.GET("/:word", func(c *gin.Context) {
		word := c.Param("word")
		res, err := Get(word)
		if err != nil {
			log.Println(err)
		}

		var data []string

		if err = json.Unmarshal(res, &data); err != nil {
			log.Printf("Error occured while unmarshalling: %v", err)
		}

		c.JSON(http.StatusOK, gin.H{
			"definition": data,
		})
	})

	return r
}
