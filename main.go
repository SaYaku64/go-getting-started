package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	v1 := router.Group("/api/v1")
	{
		v1.GET("/rick", rickHandler)
		v1.GET("/say/:name", sayHandler)
		v1.POST("/calculate", calculateHandler)
	}

	router.Run(":" + port)
}

func rickHandler(c *gin.Context) {
	c.String(http.StatusOK, "https://youtu.be/dQw4w9WgXcQ?si=ex-7DZgrTus1Vu8K")
}

func sayHandler(c *gin.Context) {
	name := c.Param("name")
	str := fmt.Sprintf("- Say my name.\n - %s.\n - You goddamn right!", name)

	c.String(http.StatusOK, str)
}

func calculateHandler(c *gin.Context) {
	var json struct {
		FirstNum  int    `json:"first_num"`
		SecondNum int    `json:"second_num"`
		Action    string `json:"action"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch strings.ToLower(json.Action) {
	case "addition":
		c.JSON(http.StatusOK, gin.H{
			"message": json.Action,
			"result":  json.FirstNum + json.SecondNum,
		})

		return
	case "subtraction":
		c.JSON(http.StatusOK, gin.H{
			"message": json.Action,
			"result":  json.FirstNum - json.SecondNum,
		})

		return
	case "multiplication":
		c.JSON(http.StatusOK, gin.H{
			"message": json.Action,
			"result":  json.FirstNum * json.SecondNum,
		})

		return
	case "division":
		c.JSON(http.StatusOK, gin.H{
			"message": json.Action,
			"result":  json.FirstNum / json.SecondNum,
		})

		return
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": json.Action,
		"result":  "Unknown action",
	})
}
