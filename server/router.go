package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func initRouter() *gin.Engine{
	r := gin.New()

	r.Use(gin.Recovery())

	r.POST("/:word", func(c *gin.Context) {
		word := c.Param("word")

		res,err := Get(word)
		if err != nil {
			log.Println(err)
		}

		log.Println(res)
	})
}